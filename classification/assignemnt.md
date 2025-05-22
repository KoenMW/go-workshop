# classification assignment 

## Objective:
Your assignment will be to check wich of the following weapon stats is most likely to be a Glintstone staff

## Weapon Stats to Evaluate:
1. Stats: 68, 0, 0, 0, 0, 4, 9, 20, 15, 0, 0, 0, 0, 0, 6, 12
2. Stats: 10, 5, 0, 3, 0, 1, 2, 4, 6, 0, 1, 0, 0, 0, 1, 5
3. Stats: 35, 20, 0, 10, 0, 7, 6, 18, 14, 5, 4, 2, 0, 1, 3, 15
4. Stats: 5, 3, 0, 1, 0, 0, 1, 2, 3, 40, 45, 50, 20, 0, 0, 3
5. Stats: 60, 15, 5, 5, 0, 6, 10, 30, 25, 10, 0, 0, 0, 0, 5, 25
6. Stats: 15, 5, 2, 8, 4, 3, 5, 10, 12, 8, 12, 15, 10, 3, 4, 10


## Check out the files 
Try to understand what is happening in classifier.go and predict.go

## Task: Write your `main.go` code

1. First Initialize the db
    classification.Init(5, "data/elden_ring_weapon.csv")

2. Then predict the class 
	If this step doesnt work see hint underneath 

3. Than print the result 
	fmt.Println("Predicted weapon type:", prediction)












	
































Use 
	prediction := classification.SOMETHING 