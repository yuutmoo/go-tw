package twitter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"net/http"
	"net/url"
	"strings"
	"time"
	"xscript/backend/go-log"
	"xscript/backend/utils"
)

var defaultHeader = map[string]string{
	"accept":             "*/*",
	"accept-language":    "zh-CN,zh;q=0.9",
	"content-type":       "application/json",
	"sec-ch-ua":          `"Chromium";v="136", "Google Chrome";v="136", "Not.A/Brand";v="99"`,
	"sec-ch-ua-mobile":   "?0",
	"sec-ch-ua-platform": `"Windows"`,
	"sec-fetch-dest":     "empty",
	"sec-fetch-mode":     "cors",
	"Origin":             "https://x.com",
	"sec-fetch-site":     "same-site",
	"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Edg/136.0.0.0",
}

type Twitter struct {
	TwName    string
	authToken string

	client *req.Client
}

func NewTwitter(twName, authToken, proxy string) *Twitter {
	cookies := []*http.Cookie{
		{Name: "auth_token", Value: authToken},
	}

	c := req.NewClient().
		SetCommonRetryCount(1).
		SetCommonRetryFixedInterval(3*time.Second).
		SetCommonRetryCondition(func(resp *req.Response, err error) bool {
			return err != nil || resp.StatusCode == http.StatusTooManyRequests
		}).
		SetCommonCookies(cookies...).
		SetCommonHeaders(defaultHeader).
		SetProxyURL(proxy).
		SetCommonBearerAuthToken("AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA").
		SetCommonHeader("x-twitter-active-user", "yes").
		SetCommonHeader("x-twitter-auth-type", "OAuth2Session").
		SetCommonHeader("x-twitter-client-language", "en")
	resp, err := c.R().Get("https://x.com")
	if err != nil {
		log.Error(err)
		return nil
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "ct0" {
			c.SetCommonHeader("x-csrf-token", cookie.Value)
		}
	}
	c.SetCommonCookies(resp.Cookies()...)

	return &Twitter{twName, authToken, c}
}
func (t *Twitter) Like(tweetURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return t.LikeWithContext(ctx, tweetURL)
}

func (t *Twitter) LikeWithContext(ctx context.Context, tweetURL string) error {
	split := strings.Split(tweetURL, "/")
	if len(split) < 2 {
		return errors.New("wrong tweet url")
	}

	tweetId := split[len(split)-1]
	api := "https://x.com/i/api/graphql/lI07N6Otwv1PhnEgXILM7A/FavoriteTweet"
	likeReq := &LikeRequest{
		QueryId: "lI07N6Otwv1PhnEgXILM7A",
		Variables: struct {
			TweetId string `json:"tweet_id"`
		}{TweetId: tweetId},
	}
	resp, err := t.client.R().
		SetContext(ctx).
		SetHeader("x-client-transaction-id", utils.GenerateClientTransactionID()).
		SetBody(likeReq).Post(api)
	log.Debug(resp)
	if err != nil {
		return err
	}
	var likeResp LikeResponse
	err = resp.Into(&likeResp)
	if err != nil {
		return err
	}
	if likeResp.Data.FavoriteTweet == "Done" {
		return nil
	}
	if len(likeResp.Errors) > 0 {
		return errors.New(likeResp.Errors[0].Message)
	}
	return errors.New("liked failed")
}

func (t *Twitter) Follow(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return t.FollowWithContext(ctx, name)
}

func (t *Twitter) FollowWithContext(ctx context.Context, name string) error {
	userId, err := t.GetUserIdByName(name)
	if err != nil {
		return err
	}
	api := "https://x.com/i/api/1.1/friendships/create.json"
	data := fmt.Sprintf("include_profile_interstitial_type=1&include_blocking=1&include_blocked_by=1&include_followed_by=1&include_want_retweets=1&include_mute_edge=1&include_can_dm=1&include_can_media_tag=1&include_ext_is_blue_verified=1&include_ext_verified_type=1&include_ext_profile_image_shape=1&skip_status=1&user_id=%s", userId)
	resp, err := t.client.R().
		SetHeader("x-client-transaction-id", utils.GenerateClientTransactionID()).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetContext(ctx).
		SetBody(data).Post(api)
	if err != nil {
		return err
	}
	log.Debug(resp)
	var followResp FollowResponse
	err = resp.Into(&followResp)
	if err != nil {
		return err
	}
	if len(followResp.Errors) > 0 {
		return errors.New(followResp.Errors[0].Message)
	}

	if resp.StatusCode != 200 {
		return errors.New("follow Failed")
	}
	return nil
}

func (t *Twitter) GetUserIdByName(screenName string) (string, error) {
	v := url.QueryEscape(fmt.Sprintf("{\"screen_name\":\"%s\"}", screenName))
	u := fmt.Sprintf("https://x.com/i/api/graphql/1sAf0uU4-B2ZLJGUX5O7LQ/ProfileSpotlightsQuery?variables=%s", v)
	resp, err := t.client.R().
		Get(u)
	if err != nil {
		return "", err
	}
	var profile ProfileResponse
	resp.Into(&profile)

	return profile.Data.UserResultByScreenName.Result.RestId, nil
}

func GetFeatures() Features {
	jsonData := []byte(`{
  "premium_content_api_read_enabled": false,
  "communities_web_enable_tweet_community_results_fetch": true,
  "c9s_tweet_anatomy_moderator_badge_enabled": true,
  "responsive_web_grok_analyze_button_fetch_trends_enabled": false,
  "responsive_web_grok_analyze_post_followups_enabled": true,
  "responsive_web_jetfuel_frame": true,
  "responsive_web_grok_share_attachment_enabled": true,
  "responsive_web_edit_tweet_api_enabled": true,
  "graphql_is_translatable_rweb_tweet_is_translatable_enabled": true,
  "view_counts_everywhere_api_enabled": true,
  "longform_notetweets_consumption_enabled": true,
  "responsive_web_twitter_article_tweet_consumption_enabled": true,
  "tweet_awards_web_tipping_enabled": false,
  "responsive_web_grok_show_grok_translated_post": false,
  "responsive_web_grok_analysis_button_from_backend": false,
  "creator_subscriptions_quote_tweet_preview_enabled": false,
  "longform_notetweets_rich_text_read_enabled": true,
  "longform_notetweets_inline_media_enabled": true,
  "payments_enabled": false,
  "profile_label_improvements_pcf_label_in_post_enabled": true,
  "rweb_tipjar_consumption_enabled": true,
  "verified_phone_label_enabled": false,
  "articles_preview_enabled": true,
  "responsive_web_grok_community_note_auto_translation_is_enabled": false,
  "responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
  "freedom_of_speech_not_reach_fetch_enabled": true,
  "standardized_nudges_misinfo": true,
  "tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": true,
  "responsive_web_grok_image_annotation_enabled": true,
  "responsive_web_graphql_timeline_navigation_enabled": true,
  "responsive_web_enhance_cards_enabled": false
}`)
	var features Features
	err := json.Unmarshal(jsonData, &features)
	if err != nil {
		log.Error(err)
	}
	return features

}

func (t *Twitter) Comment(tweetURL, text string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return t.CommentWithContext(ctx, tweetURL, text)
}

func (t *Twitter) CommentWithContext(ctx context.Context, tweetURL, text string) error {
	split := strings.Split(tweetURL, "/")
	if len(split) < 2 {
		return errors.New("wrong tweet url")
	}

	tweetId := split[len(split)-1]
	api := "https://x.com/i/api/graphql/SoVnbfCycZ7fERGCwpZkYA/CreateTweet"
	commentReq := &CommentRequest{
		QueryId: "SoVnbfCycZ7fERGCwpZkYA",
		Variables: struct {
			TweetText string `json:"tweet_text"`
			Reply     struct {
				InReplyToTweetId    string        `json:"in_reply_to_tweet_id"`
				ExcludeReplyUserIds []interface{} `json:"exclude_reply_user_ids"`
			} `json:"reply"`
			DarkRequest bool `json:"dark_request"`
			Media       struct {
				MediaEntities     []interface{} `json:"media_entities"`
				PossiblySensitive bool          `json:"possibly_sensitive"`
			} `json:"media"`
			SemanticAnnotationIds  []interface{} `json:"semantic_annotation_ids"`
			DisallowedReplyOptions interface{}   `json:"disallowed_reply_options"`
		}{
			TweetText: text,
			Reply: struct {
				InReplyToTweetId    string        `json:"in_reply_to_tweet_id"`
				ExcludeReplyUserIds []interface{} `json:"exclude_reply_user_ids"`
			}{InReplyToTweetId: tweetId, ExcludeReplyUserIds: []interface{}{}},
			DarkRequest: false,
			Media: struct {
				MediaEntities     []interface{} `json:"media_entities"`
				PossiblySensitive bool          `json:"possibly_sensitive"`
			}{MediaEntities: []interface{}{}, PossiblySensitive: false},
			SemanticAnnotationIds:  []interface{}{},
			DisallowedReplyOptions: nil,
		},
		Features: GetFeatures(),
	}
	resp, err := t.client.R().
		SetContext(ctx).
		SetHeader("x-client-transaction-id", utils.GenerateClientTransactionID()).
		SetBody(commentReq).Post(api)
	log.Debug(resp)
	if err != nil {
		return err
	}
	var commentResp CommentResponse
	err = resp.Into(&commentResp)
	if err != nil {
		return err
	}
	if len(commentResp.Errors) > 0 {
		return errors.New(commentResp.Errors[0].Message)
	}
	if commentResp.Data.CreateTweet.TweetResults.Result.RestId != "" {
		return nil
	}
	return errors.New("comment failed")
}

func (t *Twitter) Retweet(tweetURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return t.RetweetWithContext(ctx, tweetURL)
}

func (t *Twitter) RetweetWithContext(ctx context.Context, tweetURL string) error {
	split := strings.Split(tweetURL, "/")
	if len(split) < 2 {
		return errors.New("wrong tweet url")
	}

	tweetId := split[len(split)-1]
	api := "https://x.com/i/api/graphql/ojPdsZsimiJrUGLR1sjUtA/CreateRetweet"
	retweetReq := &RetweetRequest{
		QueryId: "ojPdsZsimiJrUGLR1sjUtA",
		Variables: struct {
			TweetId     string `json:"tweet_id"`
			DarkRequest bool   `json:"dark_request"`
		}{TweetId: tweetId, DarkRequest: false},
	}
	resp, err := t.client.R().
		SetContext(ctx).
		SetHeader("x-client-transaction-id", utils.GenerateClientTransactionID()).
		SetBody(retweetReq).Post(api)
	log.Debug(resp)
	if err != nil {
		return err
	}
	var retweetResp RetweetResponse
	err = resp.Into(&retweetResp)
	if err != nil {
		log.Debug(resp,resp.StatusCode)
		return err
	}
	if len(retweetResp.Errors) > 0 {
		return errors.New(retweetResp.Errors[0].Message)
	}

	if retweetResp.Data.CreateRetweet.RetweetResults.Result.RestId != "" {
		return nil
	}
	return errors.New("retweet failed")
}
