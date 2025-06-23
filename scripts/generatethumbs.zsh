mkdir -p ./thumbs;

for file in ./*.{jpg,JPG,png,webp,WEBP}; do
  magick "$file" -resize 800x800 "./thumbs/$(basename "$file")"
done