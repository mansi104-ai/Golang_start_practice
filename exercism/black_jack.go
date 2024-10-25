package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
    case "ace":
        return 11
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4
    case "five":
        return 5
    case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    case "ten", "jack", "queen", "king":
        return 10
    default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// panic("Please implement the FirstTurn function")
    // Card1 = ParseCard(card1)
    // Card2 = ParseCard(card2)
    // Dealer = ParseCard(dealerCard)
    // Figure := "ten"|| "jack" || "queen"||"king"

    // switch card1, card2{
    //     case card1 == "ace" && card2 == "ace" :
    //     return "P"

    //     case ParseCard(card1)+ ParseCard(card2) == 21:
    //     if dealerCard != Figure || "ace"{
    //         return "W"
    //     }else{
    //         return "S"
    //     }
        
    //     case 17<= ParseCard(card1)+ ParseCard(card2) && ParseCard(card1)+ ParseCard(card2) <= 20:
    //     return "S"

    //     case 12<= ParseCard(card1)+ ParseCard(card2) && ParseCard(card1)+ ParseCard(card2) <= 16 :
    //     if dealerCard < 7{
    //         return "S"
    //     }else{
    //         return "H"
    //     } 

    //     case ParseCard(card1)+ ParseCard(card2) <= 11:
    //     return "H"

        
    // }

    if card1 == "ace" && card2 == "ace"{
        return "P"
    }
    if ParseCard(card1) + ParseCard(card2) == 21 {
        if ParseCard(dealerCard) != 11 &&ParseCard(dealerCard) != 10{
            return "W"
        }else{
            return "S"
        }
        
    }
    if  17<= ParseCard(card1) + ParseCard(card2) && ParseCard(card1) + ParseCard(card2) <= 20 {
        return "S"
    }
    if 12<=ParseCard(card1) + ParseCard(card2) && ParseCard(card1) + ParseCard(card2) <= 16 {
        if ParseCard(dealerCard)< 7{
            return "S"
        }else{
            return "H"
        }
    }

    if ParseCard(card1) + ParseCard(card2) <= 11{
        return "H"
    }
    return ""

    

    
}