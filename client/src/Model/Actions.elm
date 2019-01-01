module Model.Actions exposing (Action, Actions, actionDecoder, actionsDecoder, init)

import Json.Decode as JD
import Json.Decode.Pipeline as JP


type alias Claim =
    { id : Int
    , userId : Int
    , userName : String
    , isApproved : Bool
    , text : String
    , logoFile : String
    }


type alias Action =
    { id : Int
    , purpose : String
    , name : String
    , symbol : String
    , totalSupply : String
    , claims : List Claim
    , creatorId : Int
    , creatorName : String
    , favouriteCount : Int
    , didUserLike : Bool
    }


type alias Actions =
    { entries : List Action
    }


init : Actions
init =
    { entries = []
    }


actionsDecoder : JD.Decoder Actions
actionsDecoder =
    JP.decode Actions
        |> JP.required "Entries" (JD.list actionDecoder)


actionDecoder : JD.Decoder Action
actionDecoder =
    JP.decode Action
        |> JP.required "ID" JD.int
        |> JP.required "Purpose" JD.string
        |> JP.required "Name" JD.string
        |> JP.required "Symbol" JD.string
        |> JP.required "TotalSupply" JD.string
        |> JP.required "Claims" (JD.list claimDecoder)
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
