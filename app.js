document.addEventListener("DOMContentLoaded", () => {
  navigateToPage(location.pathname, false);

  document.querySelectorAll("button").forEach((button) => {
    button.addEventListener("click", (e) => {
      console.log("button clicked");
      e.preventDefault();
      const pageId = e.target.id.replace("goto ", "");
      navigateToPage("/" + pageId, true);
    });
  });

  window.addEventListener("popstate", (e) => {
    navigateToPage(location.pathname, false);
  });
});

function navigateToPage(pathname, updateHistory) {
  const pageId = pathname === "/" ? "1" : pathname.substring(1);
  if (pageId !== "1" && pageId !== "2" && pageId !== "3") {
    document.getElementById("error").style.display = "block";
    return;
  }
  document.getElementById("error").style.display = "none";
  document.querySelectorAll("div").forEach((div) => {
    div.style.display = "none";
  });
  document.getElementById(pageId).style.display = "block";
  if (updateHistory) {
    // 2nd argument doesn't do anything. It exists for historical reasons, but must be included. You can pass any value and it will be ignored.
    history.pushState({ pageId }, "", `/${pageId}`);
  }
}
