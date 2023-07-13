package quotes

import (
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
)

// quoteDB is a quote database. It is a slice of quotes stored in memory.
// Generally, you would use a database to store data, but for the sake of simplicity, we use a slice.
//
// For the purpose of this example, I am using a list of quotes I personally live by.
var quoteDB = []model.Quote{
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "How you do anything is how you do everything.",
		Author: "T. Harv Eker",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "Make the best use of what is in your power and take the rest as it happens.",
		Author: "Epictetus",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "The impediment to action advances action. What stands in the way becomes the way.",
		Author: "Marcus Aurelius",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "The best revenge is not to be like your enemy.",
		Author: "Marcus Aurelius",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "True happiness is to enjoy the present, without anxious dependence upon the future.",
		Author: "Seneca",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "The happiness of your life depends upon the quality of your thoughts.",
		Author: "Marcus Aurelius",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "It's not what happens to you, but how you react to it that matters.",
		Author: "Epictetus",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "If it is not right do not do it; if it is not true do not say it.",
		Author: "Marcus Aurelius",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "You have power over your mind, not outside events. Realize this and you will find strength.",
		Author: "Marcus Aurelius",
	},
	{
		ID: ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text: "Live a good life. If there are gods and they are just, then they will not care how devout you have been," +
			" but will welcome you based on the virtues you have lived by. If there are gods, but unjust, then you" +
			" should not want to worship them. If there are no gods, then you will be gone, but will have lived " +
			"a noble life that will live on in the memories of your loved ones.",
		Author: "Marcus Aurelius",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "The only limit to our realization of tomorrow will be our doubts of today.",
		Author: "Franklin D. Roosevelt",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "Success is not final, failure is not fatal: It is the courage to continue that counts.",
		Author: "Winston S. Churchill",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "Believe you can and you're halfway there.",
		Author: "Theodore Roosevelt",
	},
	{
		ID: ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text: "The quality of a person's life is in direct proportion to their commitment to excellence, regardless " +
			"of their chosen field of endeavor.",
		Author: "Vince Lombardi",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "I attribute my success to this: I never gave or took any excuse.",
		Author: "Florence Nightingale",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "Do not let what you cannot do interfere with what you can do.",
		Author: "John Wooden",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "You miss 100% of the shots you don't take.",
		Author: "Wayne Gretzky",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "The most difficult thing is the decision to act, the rest is merely tenacity.",
		Author: "Amelia Earhart",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "Hard work beats talent when talent doesn't work hard.",
		Author: "Tim Notke",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "There are no secrets to success. It is the result of preparation, hard work, and learning from failure.",
		Author: "Colin Powell",
	},
	{
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text:   "Success is walking from failure to failure with no loss of enthusiasm.",
		Author: "Winston S. Churchill",
	},
	{
		ID: ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text: "Success is not the key to happiness. Happiness is the key to success. If you love what you are " +
			"doing, you will be successful.",
		Author: "Albert Schweitzer",
	},
	{
		ID: ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text: "Your work is going to fill a large part of your life, and the only way to be truly satisfied is to do" +
			" what you believe is great work.",
		Author: "Steve Jobs",
	},
	{
		ID: ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
		Text: "If you want to achieve excellence, you can get there today. As of this second, quit doing " +
			"less-than-excellent work.",
		Author: "Thomas J. Watson",
	},
}

// GetQuotes returns all hardcoded quotes.
// This is done to avoid referencing a global variable directly.
func GetQuotes() []model.Quote {
	return quoteDB
}
