# docgen

docgen is an opinionated static site documentation generator. It does not have infinite customizability or crazy configs. It is built for normal people. That is reflected in the setup time and speed of docgen.

## Setup

To install docgen, download a release binary from [here](https://github.com/luca-rand/docgen/releases) and place it in a directory in your path.

To set up docgen in your project:

1. Create a docgen.yml file in the root of your project.
2. Open it and enter `output_folder: public`.
3. Create a README.md file at the root of your project - this will become your homepage.
4. Create any documentation pages you want in the `docs` folder. These should also be markdown files.
5. Run `docgen`.

docgen is opinionated. It only renders out `README.md` (to /) and and markdown files in `docs/` to (docs/).

## Plugins

You can alter docgens behaviour with the builtin plugins. They hook into docgen to generate extra files or to add to or change how your existing files render.
