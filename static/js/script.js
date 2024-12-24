
document.addEventListener("DOMContentLoaded", () => {
  const breedsButton = document.querySelector("#breeds-button");
  const votingButton = document.querySelector("#voting-button");
  const breedsSection = document.querySelector("#breeds-section");
  const favsSection = document.querySelector("#favs-section");
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
    imageContainer.style.display = "none"; // Hide image container
    favsSection.style.display = "none";
    footerNav.style.display = "none"; // Hide footer (since voting buttons aren't needed in breeds section)
  });

  // Show Voting Section when Voting Button is clicked (if not already in voting)
  votingButton.addEventListener("click", (event) => {
    event.preventDefault();

    // Check if we're not already in voting section
    if (
      breedsSection.style.display !== "none" ||
      imageContainer.style.display === "none"
    ) {
      // Show the main container (image and footer)
      imageContainer.style.display = "block";
      footerNav.style.display = "flex"; // Show the footer with thumbs-up, thumbs-down, and fav buttons
      breedsSection.style.display = "none"; // Hide breeds section
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
      .map(
        (breed) => `<option value="${breed.id}">${breed.name}</option>`
      )
      .join("");
    breedSelect.value = "abys"; // Set initial breed to Abyssinian (or the ID of the breed you want as default)

    loadBreedDetails(breedSelect.value); // Load initial breed details
  };

  // Load breed details including images and information
  const loadBreedDetails = async (breedID) => {
    const response = await fetch(
      `/custom/breed_images?breed_id=${breedID}`
    );
    const images = await response.json();

    const breedDetails = document.getElementById("breed-details");
    const breedImages = document.getElementById("breed-images");

    if (images && images.length > 0) {
      const breedInfo = images[0].breeds[0] || {};

      // Update breed details section
      breedDetails.querySelector("#breed-name").textContent =
        breedInfo.name || "N/A";
      breedDetails.querySelector(
        "#breed-origin"
      ).textContent = `Origin: ${breedInfo.origin || "Unknown"}`;
      breedDetails.querySelector(
        "#breed-id"
      ).textContent = `ID: ${breedID}`;
      breedDetails.querySelector("#breed-description").textContent =
        breedInfo.description || "No description available.";

      // Handle the Wikipedia link
      const wikiLink = breedInfo.wikipedia_url || "#";
      const wikiText = breedInfo.wikipedia_url
        ? "Wikipedia"
        : "No Wikipedia Link";

      breedDetails.querySelector("#breed-wikipedia").href = wikiLink;
      breedDetails.querySelector("#breed-wikipedia").textContent =
        wikiText;

      // Display breed images for the slider
      breedImages.innerHTML = images
        .map(
          (img, index) =>
            `<img src="${img.url}" alt="${breedInfo.name}" class="${
              index === 0 ? "active" : ""
            }">`
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

document.addEventListener("DOMContentLoaded", () => {
  const thumbsUpButton = document.querySelector(".thumbs-up");
  const thumbsDownButton = document.querySelector(".thumbs-down");
  const favsDownButton = document.querySelector(".favs-down"); // Favs button
  const favsSection = document.querySelector("#favs-section");
  const votingButton = document.querySelector("#voting-button");


  votingButton.addEventListener("click", (event) => {
    event.preventDefault();

    // Show the breeds section and hide other sections
    favsSection.style.display = "none";
    // imageContainer.style.display = "none"; // Hide image container
    // favsSection.style.display = "none";
    // footerNav.style.display = "none"; // Hide footer (since voting buttons aren't needed in breeds section)
  });
  const sendVote = async (imageID, value) => {
    const response = await fetch("/custom/vote", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: new URLSearchParams({ image_id: imageID, value: value }),
    });

    if (response.ok) {
      console.log("Vote submitted successfully");
    } else {
      console.error("Failed to submit vote");
    }
  };

  // Function to favourite the image
  const favouriteImage = async (imageID) => {
    const response = await fetch("/custom/favourite", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: new URLSearchParams({ image_id: imageID }),
    });

    if (response.ok) {
      console.log("Image favourited successfully");
    } else {
      console.error("Failed to favourite image");
    }
  };

  // Event listeners for voting
  thumbsUpButton.addEventListener("click", async (event) => {
    event.preventDefault();
    const imageID = document
      .querySelector(".cat-image")
      .src.split("/")
      .pop();
    await sendVote(imageID, 1);
  });

  thumbsDownButton.addEventListener("click", async (event) => {
    event.preventDefault();
    const imageID = document
      .querySelector(".cat-image")
      .src.split("/")
      .pop();
    await sendVote(imageID, -1);
  });

  // Event listener for favs-down button
  favsDownButton.addEventListener("click", async (event) => {
    event.preventDefault();
    const imageID = document
      .querySelector(".cat-image")
      .src.split("/")
      .pop(); // Get image ID from the URL
    await favouriteImage(imageID);
  });
});

document.addEventListener("DOMContentLoaded", () => {
const favsButton = document.querySelector("#favs-button");
const favsSection = document.querySelector("#favs-section");
const favsGallery = document.querySelector("#favs-gallery");
const imageContainer = document.querySelector(".image-container");
const breedsSection = document.querySelector("#breeds-section");
const footerNav = document.querySelector(".footer.nav");

// Function to load favorite images
const loadFavs = async () => {
try {
const response = await fetch("http://localhost:8080/custom/favourites");
const data = await response.json();

if (data && data.length > 0) {
  // Clear the previous gallery content
  favsGallery.innerHTML = "";

  // Display each favorite image
  data.forEach((fav) => {
    const imgElement = document.createElement("img");
    imgElement.src = fav.image.url; // Use the URL of the image
    imgElement.alt = "Favorite Cat Image";
    imgElement.classList.add("favs-image");

    // Append to the gallery container
    favsGallery.appendChild(imgElement);
  });
} else {
  favsGallery.innerHTML = "<p>No favorite images found.</p>";
}
} catch (error) {
console.error("Error loading favorite images:", error);
favsGallery.innerHTML = "<p>Failed to load favorites.</p>";
}
};

// Event listener for Favs button
favsButton.addEventListener("click", (event) => {
  const gridButton = document.querySelector(".grid-btn");
  const barButton = document.querySelector(".bar-btn");
  const favsGallery = document.querySelector("#favs-gallery");


event.preventDefault();
  // Event listener for Grid View button
  gridButton.addEventListener("click", () => {
    favsGallery.classList.remove("bar-view");
    favsGallery.classList.add("grid-view");
  });

  // Event listener for Bar View button
  barButton.addEventListener("click", () => {
    favsGallery.classList.remove("grid-view");
    favsGallery.classList.add("bar-view");
  });
  // Hide all other sections
  breedsSection.style.display = "none";
  imageContainer.style.display = "none";
  footerNav.style.display = "none";

  // Show the Favs section
  favsSection.style.display = "block";

  // Reset to Grid View when Favs button is clicked
  favsGallery.classList.remove("bar-view");
  favsGallery.classList.add("grid-view");

// Show the Favs section
favsSection.style.display = "block";

// Load favorite images
loadFavs();
});
});


document.addEventListener("DOMContentLoaded", () => {
  const gridButton = document.querySelector(".grid-btn");
  const barButton = document.querySelector(".bar-btn");
  const favsGallery = document.querySelector("#favs-gallery");

  // Event listener for Grid View button
  gridButton.addEventListener("click", () => {
    favsGallery.classList.remove("bar-view");
    favsGallery.classList.add("grid-view");
  });

  // Event listener for Bar View button
  barButton.addEventListener("click", () => {
    favsGallery.classList.remove("grid-view");
    favsGallery.classList.add("bar-view");
  });
});


