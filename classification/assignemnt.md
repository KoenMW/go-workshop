# classification assignment

## Objective:

Your assignment will be to check wich of the following weapon stats is most likely to be a light bow

## Weapon Stats to Evaluate:

1. Stats: 43, 40, 0, 0, 0, 25, 15, 15, 15, 1, 0, 5, 0, 3, 0, 3
2. Stats: 313, 93, 0, 0, 0, 56, 33, 27, 27, 3, 2, 0, 0, 8, 0, 10
3. Stats: 200, 60, 0, 0, 0, 4, 5, 0, 0, 0, 0, 0, 0, 4.5, 0, 2
4. Stats: 240, 62, 0, 0, 0, 47, 31, 31, 31, 2, 3, 0, 0, 3, 0, 3
5. Stats: 404, 0, 0, 0, 0, 75, 45, 45, 45, 1, 2, 0, 0, 18, 0, 25
6. Stats: 284, 0, 0, 183, 0, 62, 33, 33, 45, 2, 2, 1, 0, 9, 0, 20

## Check out the files

Try to understand what is happening in classifier.go and predict.go

## Task: Write your `main.go` code

1. First import classification in main
   "go-workshop/classification"
2. Then Initialize the db
   classification.Init(5, "data/elden_ring_weapon.csv")

3. Then predict the class
   (you can use chatgpt to get code to loop through the weapons)
   If you are stuck see hint underneath

4. Then print the result
   fmt.Println("Predicted weapon type:", prediction)

Use
prediction := classification.SOMETHING
