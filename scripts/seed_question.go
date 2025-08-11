package main

import (
	"log"

	"iq-go/internal/config"
	"iq-go/internal/database"
	"iq-go/internal/models"
)

func main() {
	cfg := config.Load()
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create default test
	test := models.Test{
		Name:        "Cognitive Assessment",
		Description: "A comprehensive cognitive assessment test covering analytical reasoning, working memory, processing speed, attention & focus, and emotional regulation",
		Duration:    60, // 60 minutes
	}
	db.Create(&test)

	// Seed questions
	questions := []models.Question{
		// Analytical Reasoning (1-10)
		{TestID: test.ID, QuestionText: "What is the next number in this sequence? 3, 9, 27, 81, ___", QuestionType: models.MultipleChoice, Category: models.AnalyticalReasoning, Options: `["108", "162", "243", "324"]`, CorrectAnswer: "c", OrderIndex: 1, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "Logic chains: All bloops are razzles. All razzles are squibs. Which statement must be true?", QuestionType: models.MultipleChoice, Category: models.AnalyticalReasoning, Options: `["All squibs are bloops", "All bloops are squibs", "Some squibs are bloops", "No bloops are squibs"]`, CorrectAnswer: "b", OrderIndex: 2, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "Odd-one-out: Which word does NOT belong with the others?", QuestionType: models.MultipleChoice, Category: models.AnalyticalReasoning, Options: `["Tulip", "Rose", "Oak", "Lily"]`, CorrectAnswer: "c", OrderIndex: 3, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "Analogy: Cell is to Organ as Brick is to _____.", QuestionType: models.MultipleChoice, Category: models.AnalyticalReasoning, Options: `["Cement", "Wall", "Mortar", "Clay"]`, CorrectAnswer: "b", OrderIndex: 4, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "In a standard 52-card deck, how many cards are both red and face cards?", QuestionType: models.MultipleChoice, Category: models.AnalyticalReasoning, Options: `["2", "4", "6", "8"]`, CorrectAnswer: "c", OrderIndex: 5, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "If 1 January 2026 falls on a Thursday, what weekday is 1 January 2027?", QuestionType: models.MultipleChoice, Category: models.AnalyticalReasoning, Options: `["Friday", "Saturday", "Sunday", "Monday"]`, CorrectAnswer: "a", OrderIndex: 6, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "Given A > B and B > C, which statement must be true?", QuestionType: models.MultipleChoice, Category: models.AnalyticalReasoning, Options: `["A > C", "C > A", "B > A", "A = C"]`, CorrectAnswer: "a", OrderIndex: 7, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "A bookstore sells pens at ₹15 each or a box of 5 for ₹60. What is the minimum cost to buy exactly 13 pens?", QuestionType: models.MultipleChoice, Category: models.AnalyticalReasoning, Options: `["₹150", "₹165", "₹180", "₹195"]`, CorrectAnswer: "b", OrderIndex: 8, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "Which fraction is exactly halfway between ⅓ and ½?", QuestionType: models.MultipleChoice, Category: models.AnalyticalReasoning, Options: `["5⁄12", "7⁄12", "5⁄6", "2⁄5"]`, CorrectAnswer: "a", OrderIndex: 9, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "In a three-circle Venn diagram, how many regions contain exactly two sets but not the third?", QuestionType: models.MultipleChoice, Category: models.AnalyticalReasoning, Options: `["1", "2", "3", "4"]`, CorrectAnswer: "c", OrderIndex: 10, TimeLimit: 30},

		// Working Memory (11-20)
		{TestID: test.ID, QuestionText: "Digits appear for 3 seconds: 7 2 9 4 6. Type them in the same order.", QuestionType: models.TextInput, Category: models.WorkingMemory, CorrectAnswer: "72946", OrderIndex: 11, DisplayTime: 3, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "Digits appear for 3 seconds: 3 8 1 5. Type them in reverse order.", QuestionType: models.TextInput, Category: models.WorkingMemory, CorrectAnswer: "5183", OrderIndex: 12, DisplayTime: 3, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "Item list shows for 5 seconds: apple, chair, cloud, river, gold. What was the third item? (type one word)", QuestionType: models.TextInput, Category: models.WorkingMemory, CorrectAnswer: "cloud", OrderIndex: 13, DisplayTime: 5, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "A 4×4 grid flashes for 5 seconds. Which letter was in row 2, column 3?", QuestionType: models.TextInput, Category: models.WorkingMemory, CorrectAnswer: "2", OrderIndex: 14, DisplayTime: 5, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "Watch the key positions [up → down → right → left]. Click the positions in the same order.", QuestionType: models.KeySequence, Category: models.WorkingMemory, CorrectAnswer: "up,down,right,left", OrderIndex: 15, DisplayTime: 3, TimeLimit: 15},
		{TestID: test.ID, QuestionText: "Sentence shows for 5 seconds: 'A tiny bird perched quietly on the rusty mailbox.' What was the fifth word? (type one word)", QuestionType: models.TextInput, Category: models.WorkingMemory, CorrectAnswer: "quietly", OrderIndex: 16, DisplayTime: 5, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "Solve mentally: 12 + 7 − 3 × 2 = ? You will be asked for the result later—remember it.", QuestionType: models.NumberInput, Category: models.WorkingMemory, CorrectAnswer: "13", OrderIndex: 17, TimeLimit: 15},
		{TestID: test.ID, QuestionText: "You'll see the letters L G K Q T once. Type them alphabetically.", QuestionType: models.TextInput, Category: models.WorkingMemory, CorrectAnswer: "gklqt", OrderIndex: 18, DisplayTime: 3, TimeLimit: 15},
		{TestID: test.ID, QuestionText: "After 5s, answer: Which two adjectives appeared in 'She swiftly closed the heavy wooden door behind her.'? (type two words)", QuestionType: models.TextInput, Category: models.WorkingMemory, CorrectAnswer: "heavy,wooden", OrderIndex: 19, DisplayTime: 5, TimeLimit: 15},
		{TestID: test.ID, QuestionText: "Add these numbers mentally: 24, 57, 38. Enter the total.", QuestionType: models.NumberInput, Category: models.WorkingMemory, CorrectAnswer: "119", OrderIndex: 20, TimeLimit: 15},

		// Processing Speed (21-30)
		{TestID: test.ID, QuestionText: "14 × 6 − 32 = ?", QuestionType: models.MultipleChoice, Category: models.ProcessingSpeed, Options: `["40", "52", "56", "68"]`, CorrectAnswer: "b", OrderIndex: 21, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "Solve in under 10s: (17 + 8) ÷ 5 = ?", QuestionType: models.MultipleChoice, Category: models.ProcessingSpeed, Options: `["3", "4", "5", "25"]`, CorrectAnswer: "c", OrderIndex: 22, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "Which is the smallest decimal?", QuestionType: models.MultipleChoice, Category: models.ProcessingSpeed, Options: `["0.29", "0.294", "0.298", "0.30"]`, CorrectAnswer: "a", OrderIndex: 23, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "Without a calculator, 8% of 250 = ?", QuestionType: models.MultipleChoice, Category: models.ProcessingSpeed, Options: `["18", "20", "22", "25"]`, CorrectAnswer: "b", OrderIndex: 24, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "Pick the correctly spelled word:", QuestionType: models.MultipleChoice, Category: models.ProcessingSpeed, Options: `["Occurence", "Occurrence", "Occurrance", "Occurrense"]`, CorrectAnswer: "b", OrderIndex: 25, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "Synonym of 'candid' is _____.", QuestionType: models.MultipleChoice, Category: models.ProcessingSpeed, Options: `["Frank", "Secretive", "False", "Reserved"]`, CorrectAnswer: "a", OrderIndex: 26, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "Which pair sums to 47?", QuestionType: models.MultipleChoice, Category: models.ProcessingSpeed, Options: `["19 & 29", "23 & 24", "21 & 28", "25 & 21"]`, CorrectAnswer: "b", OrderIndex: 27, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "A train covers 90 km in 1h 30m. Its average speed is ____ km/h.", QuestionType: models.MultipleChoice, Category: models.ProcessingSpeed, Options: `["45", "60", "75", "120"]`, CorrectAnswer: "b", OrderIndex: 28, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "Unscramble the letters N O L O D N to form a city.", QuestionType: models.MultipleChoice, Category: models.ProcessingSpeed, Options: `["LONDON", "NODLON", "LONOND", "ONDLON"]`, CorrectAnswer: "a", OrderIndex: 29, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "Exactly 15 days after Thursday is _____.", QuestionType: models.MultipleChoice, Category: models.ProcessingSpeed, Options: `["Friday", "Saturday", "Sunday", "Monday"]`, CorrectAnswer: "a", OrderIndex: 30, TimeLimit: 5},

		// Attention & Focus (31-40)
		{TestID: test.ID, QuestionText: "Count the letter 'F' in: 'Finished files are the result of years of scientific study combined with the experience of years.'", QuestionType: models.MultipleChoice, Category: models.AttentionFocus, Options: `["3", "4", "6", "7"]`, CorrectAnswer: "c", OrderIndex: 31, TimeLimit: 15},
		{TestID: test.ID, QuestionText: "In 'A B A C B C A B C', how many times does the pattern 'A B' occur?", QuestionType: models.MultipleChoice, Category: models.AttentionFocus, Options: `["2", "3", "4", "5"]`, CorrectAnswer: "a", OrderIndex: 32, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "What is the middle letter of 'G A T E W A Y'?", QuestionType: models.MultipleChoice, Category: models.AttentionFocus, Options: `["A", "E", "T", "W"]`, CorrectAnswer: "d", OrderIndex: 33, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "Which numbers in 12, 15, 18, 20, 30 are divisible by both 3 and 5?", QuestionType: models.MultipleChoice, Category: models.AttentionFocus, Options: `["15 only", "30 only", "15 and 30", "None"]`, CorrectAnswer: "c", OrderIndex: 34, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "After seeing the colors 'red, blue, green, yellow, red, green, blue, red', which color appeared most?", QuestionType: models.MultipleChoice, Category: models.AttentionFocus, Options: `["Red", "Blue", "Green", "Yellow"]`, CorrectAnswer: "a", OrderIndex: 35, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "Identify the odd symbol: ♣ ♣ ♠ ♣ ♣", QuestionType: models.MultipleChoice, Category: models.AttentionFocus, Options: `["1st", "2nd", "3rd", "4th"]`, CorrectAnswer: "c", OrderIndex: 36, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "In the grid, how many 7s are there? 7247 | 5767 | 1378 | 9027", QuestionType: models.MultipleChoice, Category: models.AttentionFocus, Options: `["5", "6", "7", "8"]`, CorrectAnswer: "b", OrderIndex: 37, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "If you read 25 pages in 10 minutes, how many pages will you read in 26 minutes at the same speed?", QuestionType: models.MultipleChoice, Category: models.AttentionFocus, Options: `["52", "60", "65", "75"]`, CorrectAnswer: "c", OrderIndex: 38, TimeLimit: 10},
		{TestID: test.ID, QuestionText: "Spot the repeated word in 'She decided to to walk home.'", QuestionType: models.MultipleChoice, Category: models.AttentionFocus, Options: `["She", "decided", "to", "home"]`, CorrectAnswer: "c", OrderIndex: 39, TimeLimit: 5},
		{TestID: test.ID, QuestionText: "Fill the missing number so each row totals 22: 8 3 11 | 6 5 11 | 4 ? 11", QuestionType: models.MultipleChoice, Category: models.AttentionFocus, Options: `["6", "7", "8", "11"]`, CorrectAnswer: "b", OrderIndex: 40, TimeLimit: 10},

		// Emotional Regulation (41-50)
		{TestID: test.ID, QuestionText: "A colleague criticises you publicly. Your best initial response is to _____.", QuestionType: models.MultipleChoice, Category: models.EmotionalRegulation, Options: `["Defend your work on the spot", "Stay calm, thank them, and ask to discuss later", "Ignore the comment", "Complain to the manager"]`, CorrectAnswer: "b", OrderIndex: 41, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "Your project misses its deadline. What do you do first?", QuestionType: models.MultipleChoice, Category: models.EmotionalRegulation, Options: `["Analyse and share reasons with the team", "Find someone to blame", "Stay silent", "Promise weekend work without a plan"]`, CorrectAnswer: "a", OrderIndex: 42, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "When overwhelmed, which coping strategy is most effective?", QuestionType: models.MultipleChoice, Category: models.EmotionalRegulation, Options: `["Take a brief break to reset", "Vent to co-workers", "Push through with declining quality", "Scroll social media"]`, CorrectAnswer: "a", OrderIndex: 43, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "A teammate harshly attacks your idea. To maintain collaboration, you should _____.", QuestionType: models.MultipleChoice, Category: models.EmotionalRegulation, Options: `["Calmly explain your reasoning and invite their input", "Attack their ideas in return", "Avoid them", "Report them immediately"]`, CorrectAnswer: "a", OrderIndex: 44, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "You made a mistake that impacts a client. The best action is to _____.", QuestionType: models.MultipleChoice, Category: models.EmotionalRegulation, Options: `["Inform, apologise, and present a fix", "Hide the error", "Wait—it may resolve itself", "Blame external factors"]`, CorrectAnswer: "a", OrderIndex: 45, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "During a tense negotiation you feel anger rising. You should _____.", QuestionType: models.MultipleChoice, Category: models.EmotionalRegulation, Options: `["Pause and suggest a short break", "Raise your voice", "Accept any terms to end it", "Walk out"]`, CorrectAnswer: "a", OrderIndex: 46, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "Manager adds urgent work when you're at capacity. You should _____.", QuestionType: models.MultipleChoice, Category: models.EmotionalRegulation, Options: `["Negotiate priorities or resources", "Agree immediately", "Refuse outright", "Complain to peers only"]`, CorrectAnswer: "a", OrderIndex: 47, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "A close colleague is under-performing. You should _____.", QuestionType: models.MultipleChoice, Category: models.EmotionalRegulation, Options: `["Offer private support and ask how to help", "Publicly highlight mistakes", "Ignore it", "Report them with no warning"]`, CorrectAnswer: "a", OrderIndex: 48, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "Pre-presentation anxiety: the MOST effective quick fix is _____.", QuestionType: models.MultipleChoice, Category: models.EmotionalRegulation, Options: `["Two-minute deep-breathing", "Large coffee", "Rewrite slides last minute", "Cancel the talk"]`, CorrectAnswer: "a", OrderIndex: 49, TimeLimit: 30},
		{TestID: test.ID, QuestionText: "After a heated argument, the best way to restore relations is _____.", QuestionType: models.MultipleChoice, Category: models.EmotionalRegulation, Options: `["Hold a follow-up talk to clarify and plan next steps", "Pretend it never happened", "Avoid future work together", "E-mail proving you were right"]`, CorrectAnswer: "a", OrderIndex: 50, TimeLimit: 30},
	}

	for _, question := range questions {
		db.Create(&question)
	}

	log.Println("Database seeded successfully!")
}
