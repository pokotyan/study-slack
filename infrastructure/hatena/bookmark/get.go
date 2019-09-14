// はてブのAPI、ジャンルごとに記事データ取得とかができない。
// あくまでできるのはURLをこちらから指定して、その詳細情報を取得するAPIとかしかない。
// なので、スクレイピングして各ジャンルのランキングにあるURLを取って来て、そのURLを元にはてなAPIを叩いて詳細情報取得しよう。
// https://note.nkmk.me/python-scrapy-hatena-bookmark-api/
// http://developer.hatena.ne.jp/ja/documents/bookmark/apis/getinfo