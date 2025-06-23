function loadPhotos() {
    fetch('/img/')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(photos => {
            const gallery = document.getElementById("gallery");
            gallery.innerHTML = ''; // Clear existing photos
            photos.forEach(name => {
                const img = document.createElement("img");
                img.src = `./img/${name}`;
                img.className = "photo";
                gallery.appendChild(img);
            });
        })
        .catch(error => {
            console.error('There was a problem fetching the photos:', error);
        });
}

document.addEventListener("DOMContentLoaded", () => {
    loadPhotos();
})