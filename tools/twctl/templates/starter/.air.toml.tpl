root = "."
tmp_dir = "tmp"

[build]
args_bin = ["serve"]
bin = "./tmp/main"
cmd = "templ generate && go build -o ./tmp/main ."
delay = 1000
exclude_dir = ["assets", "tmp", "vendor"]
exclude_file = []
exclude_regex = [".*_templ.go"]
exclude_unchanged = false
follow_symlink = false
full_bin = ""
include_dir = []
include_ext = [
  "go",
  "tpl",
  "tmpl",
  "templ",
  "html",
  "yaml",
  "yml",
  "css",
  "sql",
  "js",
  "ts",
]
kill_delay = "0s"
log = "build-errors.log"
send_interrupt = false
stop_on_error = true

[color]
app = ""
build = "yellow"
main = "magenta"
runner = "green"
watcher = "cyan"

[log]
time = false

[misc]
clean_on_exit = false
