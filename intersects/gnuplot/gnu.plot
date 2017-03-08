filelist = system("ls *0.80.100.dat")
plot for [fname in filelist] fname using 2:3 title fname with linespoints lw 2;
set logscale x;
pause -1

