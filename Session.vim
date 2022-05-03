let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/projects/go/tg_bot
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
set shortmess=aoO
badd +24 main.go
badd +63 models.go
badd +19 text.txt
badd +21 ~/projects/go/tg_bot/counter.go
badd +208 /usr/lib/go/src/context/context.go
badd +101 /usr/lib/go/src/os/env.go
badd +1 ~/projects/inkit-next/inkit/src/models/shared.ts
badd +134 ~/.config/nvim/lua/plugins.lua
badd +10 /usr/lib/go/src/bytes/buffer.go
badd +15 /usr/lib/go/src/strings/builder.go
badd +4 ~/projects/go/tg_bot/.env
badd +151 ~/go/pkg/mod/github.com/go-telegram-bot-api/telegram-bot-api/v5@v5.5.1/types.go
badd +1 ~/projects/go/tg_bot/tg.go
badd +431 ~/go/pkg/mod/github.com/go-telegram-bot-api/telegram-bot-api/v5@v5.5.1/bot.go
argglobal
%argdel
edit ~/projects/go/tg_bot/tg.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
wincmd _ | wincmd |
vsplit
2wincmd h
wincmd w
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 91 + 136) / 273)
exe 'vert 2resize ' . ((&columns * 90 + 136) / 273)
exe 'vert 3resize ' . ((&columns * 90 + 136) / 273)
argglobal
balt models.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 87 - ((44 * winheight(0) + 25) / 51)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 87
normal! 034|
wincmd w
argglobal
if bufexists(fnamemodify("main.go", ":p")) | buffer main.go | else | edit main.go | endif
if &buftype ==# 'terminal'
  silent file main.go
endif
balt ~/go/pkg/mod/github.com/go-telegram-bot-api/telegram-bot-api/v5@v5.5.1/bot.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 20 - ((19 * winheight(0) + 25) / 51)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 20
normal! 0
wincmd w
argglobal
if bufexists(fnamemodify("models.go", ":p")) | buffer models.go | else | edit models.go | endif
if &buftype ==# 'terminal'
  silent file models.go
endif
balt main.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
silent! normal! zE
let &fdl = &fdl
let s:l = 70 - ((40 * winheight(0) + 25) / 51)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 70
normal! 031|
wincmd w
exe 'vert 1resize ' . ((&columns * 91 + 136) / 273)
exe 'vert 2resize ' . ((&columns * 90 + 136) / 273)
exe 'vert 3resize ' . ((&columns * 90 + 136) / 273)
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
let g:this_session = v:this_session
let g:this_obsession = v:this_session
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
