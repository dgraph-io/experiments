=== RUN   TestSize
Size=10 Size of delta: 91 fixed: 82
Size=100 Size of delta: 848 fixed: 803
Size=1000 Size of delta: 7950 fixed: 8003
Size=10000 Size of delta: 75387 fixed: 80004
Size=100000 Size of delta: 695563 fixed: 800004
Size=1000000 Size of delta: 6616608 fixed: 8000005
Size=10000000 Size of delta: 59716805 fixed: 80000005
--- PASS: TestSize (7.19s)

Running with math.MaxInt32 in rand.

=== RUN   TestSize
Size=10 Size of delta: 46 fixed: 82
Max delta: 415997731. Max fixed: 1773373084. Bits delta: 28.632000418513403. Bits fixed: 30.723848937777433
Size=100 Size of delta: 393 fixed: 803
Max delta: 100139705. Max fixed: 2142985402. Bits delta: 26.577438869626576. Bits fixed: 30.996974876406725
Size=1000 Size of delta: 3367 fixed: 8003
Max delta: 15492646. Max fixed: 2146782470. Bits delta: 23.88508022829214. Bits fixed: 30.999528866630722
Size=10000 Size of delta: 29276 fixed: 80004
Max delta: 2129660. Max fixed: 2147432911. Bits delta: 21.02219169204418. Bits fixed: 30.99996591411391
Size=100000 Size of delta: 246228 fixed: 800004
Max delta: 250875. Max fixed: 2147451756. Bits delta: 17.936609186024704. Bits fixed: 30.99997857456473
Size=1000000 Size of delta: 1942843 fixed: 8000005
Max delta: 27425. Max fixed: 2147481702. Bits delta: 14.743204000177917. Bits fixed: 30.999998692662537
Size=10000000 Size of delta: 15526157 fixed: 80000005
Max delta: 3649. Max fixed: 2147483470. Bits delta: 11.833285435584482. Bits fixed: 30.999999880418308
--- PASS: TestSize (7.23s)

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=1:-4         	 5000000	       239 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=1-4          	30000000	        41.4 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=1-4          	30000000	        51.2 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=10:-4        	 3000000	       465 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=10-4         	10000000	       120 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=10-4         	10000000	       127 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=50:-4        	 3000000	       582 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=50-4         	 3000000	       488 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=50-4         	 5000000	       251 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=100:-4       	 2000000	       610 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=100-4        	 2000000	       745 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=100-4        	 5000000	       322 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=500:-4       	 2000000	       719 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=500-4        	  500000	      3601 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=500-4        	 2000000	       800 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=1000:-4      	 2000000	       748 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=1000-4       	  200000	      6294 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=1000-4       	 1000000	      1344 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=10000:-4     	 2000000	       943 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=10000-4      	   20000	     61492 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=10000-4      	  100000	     12271 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=100000:-4    	 1000000	      1081 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=100000-4     	    2000	    738457 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=100000-4     	    3000	    403899 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=1:-4        	 1000000	      2297 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=1-4         	 3000000	       429 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=1-4         	 3000000	       477 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=10:-4       	  300000	      4187 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=10-4        	 1000000	      1351 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=10-4        	 1000000	      1552 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=50:-4       	  300000	      5415 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=50-4        	  300000	      5352 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=50-4        	  500000	      3206 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=100:-4      	  200000	      6024 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=100-4       	  200000	      9003 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=100-4       	  500000	      3768 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=500:-4      	  200000	      7781 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=500-4       	   50000	     37501 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=500-4       	  200000	     10323 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=1000:-4     	  200000	      9045 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=1000-4      	   20000	     73605 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=1000-4      	  100000	     17443 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=10000:-4    	  100000	     12418 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=10000-4     	    2000	    791586 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=10000-4     	    3000	    515521 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=1:-4       	   50000	     28284 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=1-4        	  300000	      4344 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=1-4        	  300000	      4839 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=10:-4      	   20000	     76787 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=10-4       	  100000	     22835 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=10-4       	  100000	     20040 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=50:-4      	   10000	    118906 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=50-4       	   30000	     55389 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=50-4       	   30000	     46639 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=100:-4     	   10000	    139950 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=100-4      	   20000	     91834 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=100-4      	   30000	     52808 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=500:-4     	   10000	    200302 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=500-4      	    3000	    425983 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=500-4      	   10000	    208620 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=1000:-4    	   10000	    231808 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=1000-4     	    2000	    836291 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=1000-4     	    3000	    515041 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.00:ratio=1:-4      	    3000	    482508 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.00:ratio=1-4       	   10000	    139761 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.00:ratio=1-4       	   10000	    146014 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.00:ratio=10:-4     	    2000	    961117 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.00:ratio=10-4      	    5000	    273067 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.00:ratio=10-4      	    5000	    297290 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.00:ratio=50:-4     	    1000	   1567806 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.00:ratio=50-4      	    2000	    606763 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.00:ratio=50-4      	    2000	    573698 ns/op


====================
Using delta encoding: Merge below 500, binary >= 500 ratio.
BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=1:-4         	 5000000	       241 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=1-4          	30000000	        51.7 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=1-4          	30000000	        54.6 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=10:-4        	 3000000	       464 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=10-4         	10000000	       134 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=10-4         	10000000	       153 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=50:-4        	 3000000	       550 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=50-4         	 3000000	       449 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=50-4         	 5000000	       320 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=100:-4       	 3000000	       588 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=100-4        	 2000000	       782 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=100-4        	 3000000	       401 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=500:-4       	 2000000	       724 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=500-4        	  500000	      2489 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=500-4        	 2000000	       698 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=1000:-4      	 2000000	       750 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=1000-4       	  200000	      7096 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=1000-4       	 2000000	       741 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=10000:-4     	 2000000	       897 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=10000-4      	   20000	     64910 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=10000-4      	 2000000	       867 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.00:ratio=100000:-4    	 1000000	      1044 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.00:ratio=100000-4     	    2000	    596459 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.00:ratio=100000-4     	 1000000	      1092 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=1:-4        	 1000000	      2249 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=1-4         	 3000000	       420 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=1-4         	 3000000	       489 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=10:-4       	  300000	      4120 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=10-4        	 1000000	      1271 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=10-4        	 1000000	      1413 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=50:-4       	  300000	      5200 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=50-4        	  300000	      5338 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=50-4        	  500000	      2773 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=100:-4      	  200000	      6124 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=100-4       	  200000	      8781 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=100-4       	  500000	      3656 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=500:-4      	  200000	      7654 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=500-4       	   50000	     38061 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=500-4       	  200000	     10203 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=1000:-4     	  200000	      8448 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=1000-4      	   20000	     72813 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=1000-4      	  200000	     11188 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.00:ratio=10000:-4    	  100000	     12560 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.00:ratio=10000-4     	    2000	    909598 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.00:ratio=10000-4     	  100000	     15800 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=1:-4       	   50000	     29578 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=1-4        	  300000	      4427 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=1-4        	  300000	      5358 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=10:-4      	   20000	     76606 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=10-4       	  100000	     23266 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=10-4       	  100000	     21067 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=50:-4      	   10000	    114187 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=50-4       	   30000	     55419 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=50-4       	   30000	     41590 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=100:-4     	   10000	    136911 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=100-4      	   20000	     92353 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=100-4      	   30000	     51978 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=500:-4     	   10000	    199200 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=500-4      	    3000	    422493 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=500-4      	   10000	    195395 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.00:ratio=1000:-4    	   10000	    231203 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.00:ratio=1000-4     	    2000	    883644 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.00:ratio=1000-4     	   10000	    219300 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.00:ratio=1:-4      	    3000	    475883 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.00:ratio=1-4       	   10000	    140307 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.00:ratio=1-4       	   10000	    145726 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.00:ratio=10:-4     	    2000	    950890 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.00:ratio=10-4      	    5000	    269034 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.00:ratio=10-4      	    5000	    296049 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.00:ratio=50:-4     	    1000	   1548709 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.00:ratio=50-4      	    2000	    607286 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.00:ratio=50-4      	    3000	    568551 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.00:ratio=100:-4    	    1000	   1998641 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.00:ratio=100-4     	    2000	   1015418 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.00:ratio=100-4     	    2000	    853913 ns/op


BenchmarkListIntersect/:Bin:size=10:overlap=0.80:ratio=1:-4         	 5000000	       315 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.80:ratio=1-4          	50000000	        34.2 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.80:ratio=1-4          	30000000	        56.0 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.80:ratio=10:-4        	 3000000	       451 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.80:ratio=10-4         	20000000	       102 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.80:ratio=10-4         	10000000	       161 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.80:ratio=50:-4        	 3000000	       535 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.80:ratio=50-4         	 5000000	       394 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.80:ratio=50-4         	10000000	       220 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.80:ratio=100:-4       	 3000000	       592 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.80:ratio=100-4        	 2000000	       842 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.80:ratio=100-4        	 5000000	       386 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.80:ratio=500:-4       	 2000000	       693 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.80:ratio=500-4        	  500000	      3668 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.80:ratio=500-4        	 2000000	       707 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.80:ratio=1000:-4      	 2000000	       775 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.80:ratio=1000-4       	  200000	      7118 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.80:ratio=1000-4       	 2000000	       760 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.80:ratio=10000:-4     	 2000000	       896 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.80:ratio=10000-4      	   20000	     62872 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.80:ratio=10000-4      	 2000000	       898 ns/op

BenchmarkListIntersect/:Bin:size=10:overlap=0.80:ratio=100000:-4    	 1000000	      1056 ns/op
BenchmarkListIntersect/:Mer:size=10:overlap=0.80:ratio=100000-4     	    2000	    710901 ns/op
BenchmarkListIntersect/:Two:size=10:overlap=0.80:ratio=100000-4     	 1000000	      1165 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.80:ratio=1:-4        	  500000	      2648 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.80:ratio=1-4         	 5000000	       310 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.80:ratio=1-4         	 3000000	       519 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.80:ratio=10:-4       	  300000	      4233 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.80:ratio=10-4        	 1000000	      1195 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.80:ratio=10-4        	 1000000	      1561 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.80:ratio=50:-4       	  300000	      5379 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.80:ratio=50-4        	  300000	      5072 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.80:ratio=50-4        	  500000	      2925 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.80:ratio=100:-4      	  200000	      6056 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.80:ratio=100-4       	  200000	      8900 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.80:ratio=100-4       	  500000	      3868 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.80:ratio=500:-4      	  200000	      7466 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.80:ratio=500-4       	   50000	     37530 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.80:ratio=500-4       	  200000	      9678 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.80:ratio=1000:-4     	  200000	      8021 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.80:ratio=1000-4      	   20000	     72897 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.80:ratio=1000-4      	  200000	     10876 ns/op

BenchmarkListIntersect/:Bin:size=100:overlap=0.80:ratio=10000:-4    	  200000	     11394 ns/op
BenchmarkListIntersect/:Mer:size=100:overlap=0.80:ratio=10000-4     	    2000	    789382 ns/op
BenchmarkListIntersect/:Two:size=100:overlap=0.80:ratio=10000-4     	  100000	     14822 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.80:ratio=1:-4       	   50000	     27507 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.80:ratio=1-4        	  500000	      3190 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.80:ratio=1-4        	  300000	      5218 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.80:ratio=10:-4      	   20000	     77375 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.80:ratio=10-4       	  100000	     19514 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.80:ratio=10-4       	  100000	     21548 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.80:ratio=50:-4      	   10000	    116042 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.80:ratio=50-4       	   30000	     55583 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.80:ratio=50-4       	   30000	     41597 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.80:ratio=100:-4     	   10000	    138189 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.80:ratio=100-4      	   20000	     91345 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.80:ratio=100-4      	   30000	     51005 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.80:ratio=500:-4     	   10000	    195641 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.80:ratio=500-4      	    3000	    421975 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.80:ratio=500-4      	   10000	    203381 ns/op

BenchmarkListIntersect/:Bin:size=1000:overlap=0.80:ratio=1000:-4    	   10000	    231351 ns/op
BenchmarkListIntersect/:Mer:size=1000:overlap=0.80:ratio=1000-4     	    2000	    831060 ns/op
BenchmarkListIntersect/:Two:size=1000:overlap=0.80:ratio=1000-4     	   10000	    221602 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.80:ratio=1:-4      	    3000	    425856 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.80:ratio=1-4       	   20000	     74722 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.80:ratio=1-4       	   20000	     92743 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.80:ratio=10:-4     	    2000	    975754 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.80:ratio=10-4      	    5000	    284046 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.80:ratio=10-4      	    5000	    315248 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.80:ratio=50:-4     	    1000	   1593898 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.80:ratio=50-4      	    2000	    617119 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.80:ratio=50-4      	    2000	    583353 ns/op

BenchmarkListIntersect/:Bin:size=10000:overlap=0.80:ratio=100:-4    	    1000	   2050032 ns/op
BenchmarkListIntersect/:Mer:size=10000:overlap=0.80:ratio=100-4     	    2000	   1038451 ns/op
BenchmarkListIntersect/:Two:size=10000:overlap=0.80:ratio=100-4     	    2000	    869127 ns/op

