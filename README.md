# Gophercises
Quiz Game Create a program to run timed quizes via the command line.


This exercise is broken into two parts to help simplify the process of explaining it as well as to make it easier to solve. The second part is harder than the first, so if you get stuck feel free to move on to another problem then come back to part 2 later.


Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to problems.csv (example shown below), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.


Part 1 

Flags package, csvs, strings and files