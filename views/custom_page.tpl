<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Cat Application</title>
    <link rel="stylesheet" href="static/css/style.css" />
    <link
      href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css"
      rel="stylesheet"
    />
    <style>
      /* Container styling */
      .container {
        width: 80%;
        margin: 0 auto;
        padding: 20px;
      }

      /* Navigation Button Styling */
      .nav button {
        margin: 5px;
        padding: 10px;
        background-color: #3b8b8b;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        transition: background-color 0.3s;
      }

      .nav button:hover {
        background-color: #2a6b6b;
      }

      /* Image Section Styling */
      .image-container {
        display: block;
        text-align: center;
        margin-top: 20px;
      }

      .image-container img {
        width: 100%;
        max-width: 600px;
        height: auto;
        border-radius: 10px;
      }

      /* Breeds Section Styling */
      #breeds-section {
        display: none;
        margin-top: 20px;
      }

      #breed-select {
        margin: 10px 0;
        padding: 8px;
        font-size: 16px;
        border-radius: 5px;
        border: 1px solid #ccc;
      }

      #breed-images {
        position: relative;
        width: 200px; /* Width of the image */
        height: 200px; /* Height of the image */
        overflow: hidden; /* Hide images that overflow */
        margin: 20px auto;
        border-radius: 8px;
        border: 2px solid #ddd;
      }

      #breed-images img {
        position: absolute;
        width: 100%;
        height: 100%;
        object-fit: cover;
        opacity: 0;
        transition: opacity 2s ease-in-out;
      }

      #breed-images img.active {
        opacity: 1; /* Show the active image */
      }

      /* Footer Section Styling */
      .footer.nav {
        display: flex;
        justify-content: center;
        margin-top: 20px;
      }

      .footer.nav button {
        margin: 5px;
        padding: 12px;
        background-color: #ff8c00;
        border-radius: 50%;
        border: none;
        cursor: pointer;
        transition: background-color 0.3s;
      }

      .footer.nav button:hover {
        background-color: #e07b00;
      }

      .footer.nav button i {
        font-size: 20px;
        color: white;
      }

      /* Dropdown and Text Styling */
      #breed-details {
        margin-top: 20px;
        text-align: center;
      }

      #breed-details h2 {
        font-size: 24px;
        color: #333;
      }

      #breed-details p {
        font-size: 16px;
        color: #666;
      }

      #breed-details a {
        font-size: 16px;
        color: #007bff;
        text-decoration: none;
      }

      #breed-details a:hover {
        text-decoration: underline;
      }

    </style>
  </head>
  <body>
    <div class="container">
      <div class="nav">
        <button id="voting-button">Voting</button>
        <button href="#" id="breeds-button"><i class="fa-solid fa-magnifying-glass breeds"> Breeds</i></button>
        <button href="#"><i class="fa-regular fa-heart favs">Favs</i></button>
      </div>

      <!-- Image Section -->
      <div class="image-container">
        {{if .CatImageURL}}
        <img src="{{.CatImageURL}}" alt="Random Cat Image" class="cat-image" />
        {{else}}
        <p>No image available at the moment.</p>
        {{ end }}
      </div>

      <!-- Breeds Section -->
      <div id="breeds-section" style="display: none">
        <select id="breed-select">
          <!-- Options will be dynamically loaded here -->
        </select>

        <div id="breed-images" class="slider">
          <!-- Images will be dynamically loaded here -->
        </div>

        <div id="breed-details">
          <h2 id="breed-name">Breed Name</h2>
          <p id="breed-origin">Origin:</p>
          <p id="breed-id">ID:</p>
          <p id="breed-description">Description will appear here.</p>
          <a id="breed-wikipedia" target="_blank">Wikipedia</a>
        </div>
      </div>

      <!-- Footer Section -->
      <div class="footer nav">
        <button href="#" class="favs-down"><i class="fa-regular fa-heart"></i></button>
        <button href="#" class="thumbs-up"><i class="fa-regular fa-thumbs-up"></i></button>
        <button href="#" class="thumbs-down"><i class="fa-regular fa-thumbs-down"></i></button>
      </div>
    </div>

    <!-- JavaScript Section -->
    <script>
      document.addEventListener("DOMContentLoaded", () => {
        const breedsButton = document.querySelector("#breeds-button");
        const votingButton = document.querySelector("#voting-button");
        const breedsSection = document.querySelector("#breeds-section");
        const footerNav = document.querySelector(".footer.nav");
        const imageContainer = document.querySelector(".image-container");
        const catImageElement = document.querySelector(".cat-image");
        const favsDownButton = document.querySelector(".favs-down");
        const thumbsUpButton = document.querySelector(".thumbs-up");
        const thumbsDownButton = document.querySelector(".thumbs-down");

        // Show Breeds Section when Breeds Button is clicked
        breedsButton.addEventListener("click", (event) => {
          event.preventDefault();

          // Show the breeds section and hide other sections
          breedsSection.style.display = "block";
          imageContainer.style.display = "none";  // Hide image container
          footerNav.style.display = "none";  // Hide footer (since voting buttons aren't needed in breeds section)
        });

        // Show Voting Section when Voting Button is clicked (if not already in voting)
        votingButton.addEventListener("click", (event) => {
          event.preventDefault();

          // Check if we're not already in voting section
          if (breedsSection.style.display !== "none" || imageContainer.style.display === "none") {
            // Show the main container (image and footer)
            imageContainer.style.display = "block";
            footerNav.style.display = "flex";  // Show the footer with thumbs-up, thumbs-down, and fav buttons
            breedsSection.style.display = "none";  // Hide breeds section
          } else {
            // If already in voting, do nothing
            console.log("Already in Voting Section");
          }
        });

        // Event listeners for footer buttons (for changing cat images)
        favsDownButton.addEventListener("click", (event) => {
          event.preventDefault();
          changeCatImage();
        });

        thumbsUpButton.addEventListener("click", (event) => {
          event.preventDefault();
          changeCatImage();
        });

        thumbsDownButton.addEventListener("click", (event) => {
          event.preventDefault();
          changeCatImage();
        });

        // Function to change the cat image
        const changeCatImage = async () => {
          const response = await fetch("/custom");
          const data = await response.text();
          const parser = new DOMParser();
          const doc = parser.parseFromString(data, "text/html");
          const newImageSrc = doc.querySelector(".cat-image")?.src;

          if (newImageSrc && catImageElement) {
            catImageElement.src = newImageSrc;
          }
        };

        // Load breeds into the dropdown
        const loadBreeds = async () => {
          const response = await fetch("/custom/breeds");
          const breeds = await response.json();

          const breedSelect = document.getElementById("breed-select");
          breedSelect.innerHTML = breeds
            .map((breed) => `<option value="${breed.id}">${breed.name}</option>`)
            .join("");
          breedSelect.value = "abys"; // Set initial breed to Abyssinian (or the ID of the breed you want as default)

          loadBreedDetails(breedSelect.value); // Load initial breed details
        };

        // Load breed details including images and information
        const loadBreedDetails = async (breedID) => {
          const response = await fetch(`/custom/breed_images?breed_id=${breedID}`);
          const images = await response.json();

          const breedDetails = document.getElementById("breed-details");
          const breedImages = document.getElementById("breed-images");

          if (images && images.length > 0) {
            const breedInfo = images[0].breeds[0] || {};

            // Update breed details section
            breedDetails.querySelector("#breed-name").textContent =
              breedInfo.name || "N/A";
            breedDetails.querySelector("#breed-origin").textContent = `Origin: ${
              breedInfo.origin || "Unknown"
            }`;
            breedDetails.querySelector("#breed-id").textContent = `ID: ${breedID}`;
            breedDetails.querySelector("#breed-description").textContent =
              breedInfo.description || "No description available.";

            // Handle the Wikipedia link
            const wikiLink = breedInfo.wikipedia_url || "#";
            const wikiText = breedInfo.wikipedia_url
              ? "Wikipedia"
              : "No Wikipedia Link";

            breedDetails.querySelector("#breed-wikipedia").href = wikiLink;
            breedDetails.querySelector("#breed-wikipedia").textContent = wikiText;

            // Display breed images for the slider
            breedImages.innerHTML = images
              .map(
                (img, index) =>
                  `<img src="${img.url}" alt="${breedInfo.name}" class="${index === 0 ? "active" : ""}">`
              )
              .join("");

            startImageSlider(); // Start the slider
          } else {
            breedDetails.querySelector("#breed-wikipedia").href = "#";
            breedDetails.querySelector("#breed-wikipedia").textContent =
              "No images found for this breed";
          }
        };

        // Function to start the image slider
        const startImageSlider = () => {
          const images = document.querySelectorAll("#breed-images img");
          let currentIndex = 0;

          setInterval(() => {
            // Remove the active class from the current image
            images[currentIndex].classList.remove("active");

            // Move to the next image
            currentIndex = (currentIndex + 1) % images.length;

            // Add the active class to the next image
            images[currentIndex].classList.add("active");
          }, 3000); // Change image every 3 seconds
        };

        // Handle change in breed selection
        const breedSelect = document.getElementById("breed-select");
        breedSelect.addEventListener("change", (event) => {
          loadBreedDetails(event.target.value);
        });

        // Load breeds when showing the breeds section for the first time
        if (breedSelect.options.length === 0) {
          loadBreeds();
        }
      });
    </script>
  </body>
</html>
