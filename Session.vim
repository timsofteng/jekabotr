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
badd +1 main.go
badd +66 models.go
badd +70 ~/projects/go/tg_bot/tg.go
badd +1 ~/projects/go/tg_bot/.gitignore
badd +2 jekabot
badd +7 ~/projects/go/tg_bot/Session.vim
badd +17 ~/projects/go/tg_bot/go.mod
badd +200 ~/go/pkg/mod/github.com/go-telegram-bot-api/telegram-bot-api/v5@v5.5.1/helpers.go
badd +1 .env
badd +1 /tmp/nvimpocECa/jeka_bot-messages-List-2022-05-04\ 19∶04∶46
badd +1 ~/projects/go/tg_bot/voice_ids
badd +2 /tmp/nvimpocECa/jeka_bot-query-2022-05-04\ 19∶32∶42
badd +1 /tmp/nvimpocECa/43.dbout
badd +1 /tmp/nvimpocECa/jeka_bot-query-2022-05-04\ 19∶39∶52
badd +1 /tmp/nvimlbJfLf/jeka_bot-messages-Columns-2022-05-04\ 19∶43∶49
badd +1 /tmp/nvimlbJfLf/jeka_bot-messages-Columns-2022-05-04\ 19∶51∶14
badd +6 /tmp/nvimOOMliC/jeka_bot-query-2022-05-04\ 20∶32∶26
badd +2 /tmp/nvimOOMliC/jeka_bot-query-2022-05-04\ 20∶35∶36
badd +1 /tmp/nvimOOMliC/jeka_bot-voice-Columns-2022-05-04\ 20∶37∶08
badd +1 /tmp/nvimOOMliC/jeka_bot-voice-List-2022-05-04\ 20∶37∶18
badd +2 /tmp/nvimOOMliC/jeka_bot-query-2022-05-04\ 20∶37∶52
badd +1 /tmp/nvimOOMliC/jeka_bot-messages-List-2022-05-04\ 20∶38∶08
badd +1 /tmp/nvimOOMliC/jeka_bot-query-2022-05-04\ 20∶39∶31
badd +1 /tmp/nvimOOMliC/jeka_bot-query-2022-05-04\ 20∶39∶59
badd +2 /tmp/nvimOOMliC/jeka_bot-query-2022-05-04\ 20∶41∶46
badd +2 /tmp/nvimOOMliC/jeka_bot-query-2022-05-04\ 20∶44∶41
badd +1 /tmp/nvimOOMliC/jeka_bot-text-Columns-2022-05-04\ 20∶45∶09
argglobal
%argdel
edit main.go
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
balt ~/projects/go/tg_bot/tg.go
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
let s:l = 16 - ((15 * winheight(0) + 26) / 52)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 16
normal! 013|
wincmd w
argglobal
if bufexists(fnamemodify("~/projects/go/tg_bot/tg.go", ":p")) | buffer ~/projects/go/tg_bot/tg.go | else | edit ~/projects/go/tg_bot/tg.go | endif
if &buftype ==# 'terminal'
  silent file ~/projects/go/tg_bot/tg.go
endif
balt main.go
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
let s:l = 118 - ((51 * winheight(0) + 26) / 52)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 118
normal! 0
wincmd w
argglobal
if bufexists(fnamemodify("models.go", ":p")) | buffer models.go | else | edit models.go | endif
if &buftype ==# 'terminal'
  silent file models.go
endif
balt ~/projects/go/tg_bot/tg.go
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
let s:l = 91 - ((41 * winheight(0) + 26) / 52)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 91
normal! 034|
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
