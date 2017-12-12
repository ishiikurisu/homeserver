function ytdl_main() {
  var links = document.getElementById("links");
  var results = document.getElementById("results");

  results.innerHTML = "<h1>Resultados</h1>"
                    + "<p>"
                    + "  Você escreveu:\n"
                    + links.value
                    + "</p>";
}
