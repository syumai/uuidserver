import handlers from "./build/worker.mjs";

const result = document.getElementById("result");

document.getElementById("buttons").addEventListener("click", async (e) => {
  const button = e.target.closest("button[data-path]");
  if (!button) return;
  const res = await handlers.fetch(new Request(button.dataset.path));
  result.textContent = await res.text();
});
