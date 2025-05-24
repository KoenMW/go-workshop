# clustering assignment (with visualization)

## Objective:

Your assignment is to group weapons into clusters based on their stats and visualize them on a 2D scatterplot.

## Task: Write your `main.go` code

1. First import `clustering` and `csvutil` in main

2. Load the weapon data and check for errors

3. Generate datapoints with the following recommended fields:  
   `Phy`, `Mag`, `Fir`, `Lit`, `Hol`, `Sta`, `Str`, `Dex`

4. Run `clustering.RunKMeans(points, k)` with `k = 3`

5. Project the points to 2D using `clustering.ProjectTo2D(points)`

6. Visualize with `clustering.PlotClusters(points2D, assignments, "output.png")`

Then open the file `output.png` to see the result!

💡 Try changing the number of clusters or different fields to see what changes in the visual result.

✅ **Tip**: You are allowed to use **ChatGPT** or **GitHub Copilot** to help generate the code for your `main.go`.
