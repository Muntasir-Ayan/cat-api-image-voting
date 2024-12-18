// Fetch the images from The Cat API
fetch('https://api.thecatapi.com/v1/images/search?limit=10')
  .then(response => response.json())  // Parse the JSON response
  .then(data => {
    const imageContainer = document.getElementById('catImagesContainer');

    // Select a random image from the data array
    const randomIndex = Math.floor(Math.random() * data.length); // Random index between 0 and 9
    const cat = data[randomIndex];

    // Create the img element for the selected image
    const imgElement = document.createElement('img');
    imgElement.src = cat.url;
    imgElement.alt = 'Random Cat';
    imgElement.title = 'Click to view larger';

    // Append the image to the container
    imageContainer.appendChild(imgElement);
  })
  .catch(error => {
    console.error('Error fetching cat images:', error);
  });

// Add JavaScript functionality for nav links
document.querySelectorAll('.nav a').forEach(link => {
  link.addEventListener('click', (event) => {
    event.preventDefault();
    console.log(`Clicked on "${event.target.textContent}" link`);
    // Add your desired functionality here, like navigating to a new page, etc.
  });
});
