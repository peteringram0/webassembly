const gulp = require("gulp");

gulp.task("build", callback => {
  const asc = require("assemblyscript/bin/asc");
  asc.main([
    "main.ts",
    // "--baseDir", "./",
    "--binaryFile", "./main.wasm",
    // "--sourceMap",
    "--measure"
  ], callback);
});

gulp.task("default", ["build"]);
