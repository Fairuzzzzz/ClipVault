document.addEventListener("DOMContentLoaded", function () {
  var navLinks = document.querySelectorAll("nav a");
  for (var i = 0; i < navLinks.length; i++) {
    var link = navLinks[i];
    if (link.getAttribute("href") == window.location.pathname) {
      link.classList.add("live");
      break;
    }
  }

  const body = document.body;
  const modeToggle = document.getElementById("mode-toggle");
  const moonIcon = modeToggle.querySelector(".fa-moon");
  const sunIcon = modeToggle.querySelector(".fa-sun");

  function enableDarkMode() {
    body.classList.add("dark-mode");
    localStorage.setItem("darkMode", "enabled");
  }

  function disableDarkMode() {
    body.classList.remove("dark-mode");
    localStorage.setItem("darkMode", null);
  }

  if (localStorage.getItem("darkMode") === "enabled") {
    enableDarkMode();
  }

  modeToggle.addEventListener("click", () => {
    if (body.classList.contains("dark-mode")) {
      disableDarkMode();
    } else {
      enableDarkMode();
    }
  });
});

