<img align="right" width="120px" alt="Sales &amp; Dungeons" src="./data/preview.png">
<img width="100" alt="Sales &amp; Dungeons" src="./data/round_icon.png">

[![Discord](https://img.shields.io/discord/678654745803751579?label=discord)](https://discord.gg/5MUZEjc)

**Sales & Dungeons** — Thermal Printer as D&amp;D Utility.

With Sales & Dungeons you can create highly customizable handouts, quick reference and much more for your Dungeons and Dragons (or other PnP) Sessions.
Most Thermal Printer are small in size and can be taken with you and kept right at the gaming table. Use-cases range from printing out magic items, spells
or a letter that the group found to little character sheets of your players to use as DM note. The possibilities are nearly endless!

**Warning:** This is still rough and early version. If you want to get this working the best way is to jump on the Discord and ask for help.

![Screenshot](./data/screenshot.png)

## Features

- Works on Windows, Mac and Linux
- Extensive templating system through [Nunjucks](https://mozilla.github.io/nunjucks/)
- Various connection methods
  - Windows Direct Printing
  - Raw USB Printing
  - CUPS
  - Serial
- Import & Export templates and data sources
  - Fast access to external data sources like Open5e (instant access to SRD monsters, spells and more)
- Edit templates in your favorite editor (e.g. Visual Studio Code) and get live preview

## How It Works

<img align="left" alt="Sales &amp; Dungeons" src="./data/work_graph.svg">

**Templates:** Templates are created in HTML (and CSS) in combination with the Nunjucks templating language. You can imagine
the templates as little websites. That makes it possible to use all the nice and convenient layout options that HTML and CSS
has to offer and even include any common framework you might need (e.g. Fontawesome for Icons).

**Rendered HTML:** After creating a template you can create entries with the data you want and print them.
Nunjucks will create the rendered HTML from the data you want to print.

**Rendered Image:** Then this HTML get's converted to a image. Currently this conversion is done by Chrome via the
Chrome Debug Protocol. Although Chrome seems like a huge overkill for just HTML-To-Image conversion it's the standard solution at the
moment because it supports most of the modern HTML and CSS features.

**ESC / POS Commands:** The last step before our awesome template hits the Printer is the conversion from the rendered image
to the "draw image" command of the printer.

**Printer:** The generated command will then be sent to the printer and printed. Now your template is ready to be used!

:tada: :tada: :tada:

## Pre-Build Binaries

Check our github release page: https://github.com/BigJk/snd/releases

## Printers, Templating & Building

If you want to see what printers were already tested, which settings they need, how the templates work or how you can build Sales & Dungeons yourself please visit the [**wiki**](https://github.com/BigJk/snd/wiki).
