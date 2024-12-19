package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	friendsList := friendships[username]
	suggerFriendsSet := make(map[string]struct{})
	for _, friend := range friendsList {
		for _, friendOfFriend := range friendships[friend] {
			if friendOfFriend != username && !contains(friendsList, friendOfFriend) {
				suggerFriendsSet[friendOfFriend] = struct{}{}
			}
		}
	}
	suggerFriends := make([]string, 0, len(suggerFriendsSet))
	for friend := range suggerFriendsSet {
		suggerFriends = append(suggerFriends, friend)
	}
	if len(suggerFriends) == 0 {
		return nil
	}
	return suggerFriends
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
