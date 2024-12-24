package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/razyneko/go-react-calorie-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

func AddEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var entry models.Entry

	// Bind incoming JSON to struct
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// Validation of struct fields
	validationErr := validate.Struct(entry)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	// Set new ObjectID for the entry
	entry.ID = primitive.NewObjectID()

	// Insert the new entry into the database
	result, insertErr := entryCollection.InsertOne(ctx, entry)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create entry"})
		fmt.Println(insertErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetEntries(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var entries []bson.M

	// Retrieve all entries from the collection
	cursor, err := entryCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// Decode the cursor into entries
	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, entries)
}

func GetEntryById(c *gin.Context) {
	EntryId := c.Params.ByName("id")
	docId, _ := primitive.ObjectIDFromHex(EntryId)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var entry bson.M
	if err := entryCollection.FindOne(ctx, bson.M{"_id": docId}).Decode(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, entry)
}

func GetEntriesByIngredient(c *gin.Context) {
	ingredient := c.Params.ByName("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{"ingredients": ingredient})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println()
		return
	}
	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, entries)
}

func UpdateEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, err := primitive.ObjectIDFromHex(entryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var entry models.Entry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate struct data
	validationErr := validate.Struct(entry)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		return
	}

	// Construct the update document, set fields if they are not nil
	updateFields := bson.M{}
	if entry.Dish != nil {
		updateFields["dish"] = *entry.Dish
	}
	if entry.Fat != nil {
		updateFields["fat"] = *entry.Fat
	}
	if entry.Ingredients != nil {
		updateFields["ingredients"] = *entry.Ingredients
	}
	if entry.Calories != nil {
		updateFields["calories"] = *entry.Calories
	}

	// Perform the update
	result, err := entryCollection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{"$set": updateFields},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"modifiedCount": result.ModifiedCount})
}

func UpdateIngredient(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, err := primitive.ObjectIDFromHex(entryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entry ID"})
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Ingredient struct
	type Ingredient struct {
		Ingredients *string `json:"ingredients"`
	}
	var ingredient Ingredient

	// Bind JSON to ingredient struct
	if err := c.BindJSON(&ingredient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform the update
	result, err := entryCollection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{{Key: "$set", Value: bson.M{"ingredients": ingredient.Ingredients}}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"modifiedCount": result.ModifiedCount})
}

func DeleteEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Delete the entry from the collection
	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, result.DeletedCount)
}
