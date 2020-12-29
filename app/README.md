# KUAL extension

This folder contains a bare-bones [KUAL]() extension. The sample app can be launched over the traditional UI, in which case pressing buttons might redraw parts of the Kindle interface over it.

When launching it with the `run.sh` script, the UI will be stopped and restarted after the app exits. This fixes the redrawing issue, but it effectively reboots the Kindle when the app exists, and as such will take a long time.