module Model.Tokens exposing (Token, Tokens, init, tokenDecoder, tokensDecoder)

import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Main.User exposing (User, userDecoder)


type alias Token =
    { id : Int
    , name : String
    , symbol : String
    , totalSupply : String
    , logo : String
    , purpose : String
    , blockchainAddress : String
    , txAddress : String
    }


type alias Tokens =
    { entries : List Token
    , valueInUSD : String
    , count : Int
    , page : Int
    , limit : Int
    , people : List User
    }


init : Tokens
init =
    { entries = []
    , valueInUSD = "0.00"
    , count = 0
    , page = 0
    , limit = 0
    , people = []
    }


tokensDecoder : JD.Decoder Tokens
tokensDecoder =
    JP.decode Tokens
        |> JP.required "Entries" (JD.list tokenDecoder)
        |> JP.required "ValueInUSD" JD.string
        |> JP.required "Count" JD.int
        |> JP.required "Limit" JD.int
        |> JP.required "Page" JD.int
        |> JP.required "People" (JD.list userDecoder)


tokenDecoder : JD.Decoder Token
tokenDecoder =
    JP.decode Token
        |> JP.required "ID" JD.int
        |> JP.required "Name" JD.string
        |> JP.required "Symbol" JD.string
        |> JP.required "TotalSupply" JD.string
        |> JP.required "Logo" JD.string
        |> JP.required "Purpose" JD.string
        |> JP.required "BlockchainAddress" JD.string
        |> JP.required "TxAddress" JD.string



-- |> JP.required "Decimals" JD.int
-- |> JP.required "FavoritesCount" JD.int
-- |> JP.required "Volume24" JD.float
-- |> JP.required "Change24" JD.float
-- |> JP.required "MarketPrice" JD.float
