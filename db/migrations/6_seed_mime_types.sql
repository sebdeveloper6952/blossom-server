-- +migrate Up
INSERT INTO mime_types(extension, mime_type)
VALUES 
("n/a","application/octet-stream"),
(".xpm","image/x-xpixmap"),
(".7z","application/x-7z-compressed"),
(".zip","application/zip"),
(".xlsx","application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"),
(".docx","application/vnd.openxmlformats-officedocument.wordprocessingml.document"),
(".pptx","application/vnd.openxmlformats-officedocument.presentationml.presentation"),
(".epub","application/epub+zip"),
(".jar","application/jar"),
(".odt","application/vnd.oasis.opendocument.text"),
(".ott","application/vnd.oasis.opendocument.text-template"),
(".ods","application/vnd.oasis.opendocument.spreadsheet"),
(".ots","application/vnd.oasis.opendocument.spreadsheet-template"),
(".odp","application/vnd.oasis.opendocument.presentation"),
(".otp","application/vnd.oasis.opendocument.presentation-template"),
(".odg","application/vnd.oasis.opendocument.graphics"),
(".otg","application/vnd.oasis.opendocument.graphics-template"),
(".odf","application/vnd.oasis.opendocument.formula"),
(".odc","application/vnd.oasis.opendocument.chart"),
(".sxc","application/vnd.sun.xml.calc"),
(".pdf","application/pdf"),
(".fdf","application/vnd.fdf"),
("n/a","application/x-ole-storage"),
(".msi","application/x-ms-installer"),
(".aaf","application/octet-stream"),
(".msg","application/vnd.ms-outlook"),
(".xls","application/vnd.ms-excel"),
(".pub","application/vnd.ms-publisher"),
(".ppt","application/vnd.ms-powerpoint"),
(".doc","application/msword"),
(".ps","application/postscript"),
(".psd","image/vnd.adobe.photoshop"),
(".p7s","application/pkcs7-signature"),
(".ogg","application/ogg"),
(".oga","audio/ogg"),
(".ogv","video/ogg"),
(".png","image/png"),
(".png","image/vnd.mozilla.apng"),
(".jpg","image/jpeg"),
(".jxl","image/jxl"),
(".jp2","image/jp2"),
(".jpf","image/jpx"),
(".jpm","image/jpm"),
(".jxs","image/jxs"),
(".gif","image/gif"),
(".webp","image/webp"),
(".exe","application/vnd.microsoft.portable-executable"),
("n/a","application/x-elf"),
("n/a","application/x-object"),
("n/a","application/x-executable"),
(".so","application/x-sharedlib"),
("n/a","application/x-coredump"),
(".a","application/x-archive"),
(".deb","application/vnd.debian.binary-package"),
(".tar","application/x-tar"),
(".xar","application/x-xar"),
(".bz2","application/x-bzip2"),
(".fits","application/fits"),
(".tiff","image/tiff"),
(".bmp","image/bmp"),
(".ico","image/x-icon"),
(".mp3","audio/mpeg"),
(".flac","audio/flac"),
(".midi","audio/midi"),
(".ape","audio/ape"),
(".mpc","audio/musepack"),
(".amr","audio/amr"),
(".wav","audio/wav"),
(".aiff","audio/aiff"),
(".au","audio/basic"),
(".mpeg","video/mpeg"),
(".mov","video/quicktime"),
(".mp4","video/mp4"),
(".avif","image/avif"),
(".3gp","video/3gpp"),
(".3g2","video/3gpp2"),
(".mp4","audio/mp4"),
(".mqv","video/quicktime"),
(".m4a","audio/x-m4a"),
(".m4v","video/x-m4v"),
(".heic","image/heic"),
(".heic","image/heic-sequence"),
(".heif","image/heif"),
(".heif","image/heif-sequence"),
(".mj2","video/mj2"),
(".dvb","video/vnd.dvb.file"),
(".webm","video/webm"),
(".avi","video/x-msvideo"),
(".flv","video/x-flv"),
(".mkv","video/x-matroska"),
(".asf","video/x-ms-asf"),
(".aac","audio/aac"),
(".voc","audio/x-unknown"),
(".m3u","application/vnd.apple.mpegurl"),
(".rmvb","application/vnd.rn-realmedia-vbr"),
(".gz","application/gzip"),
(".class","application/x-java-applet"),
(".swf","application/x-shockwave-flash"),
(".crx","application/x-chrome-extension"),
(".ttf","font/ttf"),
(".woff","font/woff"),
(".woff2","font/woff2"),
(".otf","font/otf"),
(".ttc","font/collection"),
(".eot","application/vnd.ms-fontobject"),
(".wasm","application/wasm"),
(".shx","application/vnd.shx"),
(".shp","application/vnd.shp"),
(".dbf","application/x-dbf"),
(".dcm","application/dicom"),
(".rar","application/x-rar-compressed"),
(".djvu","image/vnd.djvu"),
(".mobi","application/x-mobipocket-ebook"),
(".lit","application/x-ms-reader"),
(".bpg","image/bpg"),
(".sqlite","application/vnd.sqlite3"),
(".dwg","image/vnd.dwg"),
(".nes","application/vnd.nintendo.snes.rom"),
(".lnk","application/x-ms-shortcut"),
(".macho","application/x-mach-binary"),
(".qcp","audio/qcelp"),
(".icns","image/x-icns"),
(".hdr","image/vnd.radiance"),
(".mrc","application/marc"),
(".mdb","application/x-msaccess"),
(".accdb","application/x-msaccess"),
(".zst","application/zstd"),
(".cab","application/vnd.ms-cab-compressed"),
(".rpm","application/x-rpm"),
(".xz","application/x-xz"),
(".lz","application/lzip"),
(".torrent","application/x-bittorrent"),
(".cpio","application/x-cpio"),
("n/a","application/tzif"),
(".xcf","image/x-xcf"),
(".pat","image/x-gimp-pat"),
(".gbr","image/x-gimp-gbr"),
(".glb","model/gltf-binary"),
(".cab","application/x-installshield"),
(".jxr","image/jxr"),
(".parquet","application/vnd.apache.parquet"),
(".txt","text/plain"),
(".html","text/html"),
(".svg","image/svg+xml"),
(".xml","text/xml"),
(".rss","application/rss+xml"),
(".atom","application/atom+xml"),
(".x3d","model/x3d+xml"),
(".kml","application/vnd.google-earth.kml+xml"),
(".xlf","application/x-xliff+xml"),
(".dae","model/vnd.collada+xml"),
(".gml","application/gml+xml"),
(".gpx","application/gpx+xml"),
(".tcx","application/vnd.garmin.tcx+xml"),
(".amf","application/x-amf"),
(".3mf","application/vnd.ms-package.3dmanufacturing-3dmodel+xml"),
(".xfdf","application/vnd.adobe.xfdf"),
(".owl","application/owl+xml"),
(".php","text/x-php"),
(".js","application/javascript"),
(".lua","text/x-lua"),
(".pl","text/x-perl"),
(".py","text/x-python"),
(".json","application/json"),
(".geojson","application/geo+json"),
(".har","application/json"),
(".ndjson","application/x-ndjson"),
(".rtf","text/rtf"),
(".srt","application/x-subrip"),
(".tcl","text/x-tcl"),
(".csv","text/csv"),
(".tsv","text/tab-separated-values"),
(".vcf","text/vcard"),
(".ics","text/calendar"),
(".warc","application/warc"),
(".vtt","text/vtt");

-- +migrate Down
DELETE FROM mime_types;
