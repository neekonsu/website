for file in ./img/*.ARW; do
  magick "$file" "${file%.ARW}.jpg"
done