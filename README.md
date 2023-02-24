# NMRViewer

A small object for viewing Bruker NMR spectra.

In a Bruker NMR file, the raw data is stored at "/{experiment number}/fid". The Fourier transformed data is at "/{experiment number}/pdata/1a" and the processed data is at "/{experiment number}/pdata/1r".

We plan to first read the "1r" file if it exists. Then plot and show the diagram.

