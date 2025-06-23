for file in ./*.ARW; do
  magick "$file" "${file%.ARW}.jpg"
done