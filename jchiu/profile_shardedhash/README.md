20160918_155605 is the current master.

20160918_170016 is using a simple sharded hash instead of gotomic hash.

Control takes 870s. Took about 8 mins on my machine.

Treatment (sharded hash) takes 585s. Took about 5 mins on my machine.

If you take a look at memory profile, listMapShard takes about <300M.
On the other hand, gotomic takes >500M.

It remains to verify that everything is working as desired. Need to add some
tests, but let's check the overall design first.
