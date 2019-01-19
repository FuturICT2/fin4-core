module Model.Tokens exposing (Claim, Liker, Miner, Token, Tokens, init, tokenDecoder, tokensDecoder)

import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Main.User exposing (User, userDecoder)


type alias Claim =
    { id : Int
    , userId : Int
    , userName : String
    , isApproved : Bool
    , text : String
    , logoFile : String
    }


type alias Miner =
    { userId : Int
    , userName : String
    , mined : String
    }


type alias Liker =
    { userId : Int
    , userName : String
    }


type alias Token =
    { id : Int
    , purpose : String
    , name : String
    , symbol : String
    , totalSupply : String
    , claims : List Claim
    , miners : List Miner
    , likers : List Liker
    , creatorId : Int
    , creatorName : String
    , favouriteCount : Int
    , didUserLike : Bool
    }


type alias Tokens =
    { entries : List Token
    , count : Int
    , page : Int
    , limit : Int
    }


init : Tokens
init =
    { entries = []
    , count = 0
    , page = 0
    , limit = 0
    }


tokensDecoder : JD.Decoder Tokens
tokensDecoder =
    JP.decode Tokens
        |> JP.required "Entries" (JD.list tokenDecoder)
        |> JP.required "Count" JD.int
        |> JP.required "Limit" JD.int
        |> JP.required "Page" JD.int


tokenDecoder : JD.Decoder Token
tokenDecoder =
    JP.decode Token
        |> JP.required "ID" JD.int
        |> JP.required "Purpose" JD.string
        |> JP.required "Name" JD.string
        |> JP.required "Symbol" JD.string
        |> JP.required "TotalSupply" JD.string
        |> JP.required "Claims" (JD.list claimDecoder)
        |> JP.optional "Miners" (JD.list minerDecoder) []
        |> JP.optional "Likers" (JD.list likerDecoder) []
        |> JP.required "CreatorID" JD.int
        |> JP.required "CreatorName" JD.string
        |> JP.required "FavouriteCount" JD.int
        |> JP.required "DidUserLike" JD.bool


claimDecoder : JD.Decoder Claim
claimDecoder =
    JP.decode Claim
        |> JP.required "ID" JD.int
        |> JP.required "UserID" JD.int
        |> JP.required "UserName" JD.string
        |> JP.required "IsApproved" JD.bool
        |> JP.required "Text" JD.string
        |> JP.required "LogoFile" JD.string


minerDecoder : JD.Decoder Miner
minerDecoder =
    JP.decode Miner
        |> JP.required "UserID" JD.int
        |> JP.required "UserName" JD.string
        |> JP.required "Mined" JD.string


likerDecoder : JD.Decoder Liker
likerDecoder =
    JP.decode Liker
        |> JP.required "UserID" JD.int
        |> JP.required "UserName" JD.string
