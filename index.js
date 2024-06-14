document.addEventListener("DOMContentLoaded", function (event) {
  const loadTime =
    performance.timing.domContentLoadedEventEnd -
    performance.timing.navigationStart;
  document.getElementById("loadTime").textContent = loadTime;
});

function toggleDarkMode() {
  document.documentElement.classList.toggle("dark");
  const themeSwitcherIcon = document.getElementById("themeSwitcherIcon");
  if (themeSwitcherIcon.src.includes("sun.svg")) {
    themeSwitcherIcon.src = "moon.svg";
    themeSwitcherIcon.alt = "Moon icon";
  } else {
    themeSwitcherIcon.src = "sun.svg";
    themeSwitcherIcon.alt = "Sun icon";
  }
}
