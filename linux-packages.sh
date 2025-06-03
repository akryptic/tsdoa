#!/bin/bash
set -e

# Variables
BUILD_BIN="./build/bin/tsdoa"
LINUX_PACKAGING="./linux-packaging"
DIST_DIR="./dist"
ICON="./build/appicon.png"
APPIMAGE_FS="appimage-fs"

rm -rf "$DIST_DIR" "$APPIMAGE_FS"
mkdir -p "$DIST_DIR"

# ------------------------
# Build DEB package
# ------------------------

# Prepare folder structure
rm -rf "$LINUX_PACKAGING"

# Copy binary
mkdir -p "$LINUX_PACKAGING/usr/local/bin"
cp "$BUILD_BIN" "$LINUX_PACKAGING/usr/local/bin/tsdoa"
chmod +x "$LINUX_PACKAGING/usr/local/bin/tsdoa" # Ensure the binary is executable after copying

# Copy icon
mkdir -p "$LINUX_PACKAGING/usr/share/icons/hicolor/256x256/apps"
cp "$ICON" "$LINUX_PACKAGING/usr/share/icons/hicolor/256x256/apps/tsdoa.png"

mkdir -p "$LINUX_PACKAGING/DEBIAN"

# Create control file
cat > "$LINUX_PACKAGING/DEBIAN/control" <<EOF
Package: tsdoa
Version: 0.1.1
Section: utils
Priority: optional
Architecture: amd64
Maintainer: Akryptic <155140399+akryptic@users.noreply.github.com>
Description: A powerful and elegant tasks / todos manager.
 A desktop-native app built with Wails, designed to help you
 manage tasks, steps, and substeps efficiently.
EOF


# Create .desktop file
mkdir -p "$LINUX_PACKAGING/usr/share/applications"
cat > "$LINUX_PACKAGING/usr/share/applications/tsdoa.desktop" <<EOF
[Desktop Entry]
Name=tsdoa
Exec=tsdoa
Icon=/usr/share/icons/hicolor/256x256/apps/tsdoa
Type=Application
Terminal=false
Categories=Utility
EOF

cat > "$LINUX_PACKAGING/DEBIAN/postinst" <<EOF
#!/bin/bash
set -e
gtk-update-icon-cache -f /usr/share/icons/hicolor || true
EOF

chmod +x "$LINUX_PACKAGING/DEBIAN/postinst"

# Build deb package
dpkg-deb --build "$LINUX_PACKAGING" "$DIST_DIR/tsdoa-linux-amd64.deb"

printf "\n\nDeb package built: $DIST_DIR/tsdoa-linux-amd64.deb\n\n"

# ------------------------
# Prepare AppImage directory
# ------------------------
cp "$LINUX_PACKAGING/usr/share/applications/tsdoa.desktop" "$LINUX_PACKAGING/tsdoa.desktop"
rm -rf "$LINUX_PACKAGING/DEBIAN" "$LINUX_PACKAGING/usr/share/applications"

# Download and extract appimagetool
curl -L -o appimagetool.AppImage https://github.com/AppImage/AppImageKit/releases/latest/download/appimagetool-x86_64.AppImage
chmod +x appimagetool.AppImage

# Exctract AppRun from appimagetool
./appimagetool.AppImage --appimage-extract > /dev/null
mv squashfs-root "$APPIMAGE_FS"

# Copy AppRun (make sure you have AppRun in your repo or build process)
cat > "$LINUX_PACKAGING/AppRun" <<EOF
#!/bin/bash
exec $APPDIR/usr/local/bin/tsdoa
EOF

chmod +x "$LINUX_PACKAGING/AppRun"

# ------------------------
# Build AppImage
# ------------------------

# Assuming appimagetool is installed and in PATH
chmod +x "$APPIMAGE_FS"/AppRun
$APPIMAGE_FS/AppRun --comp xz "$LINUX_PACKAGING" "$DIST_DIR/tsdoa-linux-amd64.AppImage"
printf "\n\nAppImage built: $DIST_DIR/tsdoa-linux-amd64.AppImage\n"

cp "$LINUX_PACKAGING/usr/local/bin/tsdoa" "$DIST_DIR/tsdoa-linux-amd64"