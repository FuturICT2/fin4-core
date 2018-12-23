module Model.Actions exposing (Action, Actions, actionDecoder, actionsDecoder, init)

import Json.Decode as JD
import Json.Decode.Pipeline as JP


type alias Proposal =
    { id : Int
    , description : String
    , doerId : Int
    , doerName : String
    , approved : Bool
    }


type alias Supporter =
    { userId : Int
    , userName : String
    , tokenId : Int
    , tokenName : String
    , amount : String
    , status : Int
    }


type alias Action =
    { id : Int
    , description : String
    , creatorId : Int
    , creatorName : String
    , status : Int
    , logoFile : String
    , startsAt : String
    , endsAt : String
    , proposals : List Proposal
    , supporters : List Supporter
    , endsInHours : String
    , endsInMinutes : String
    , isTimeLimit : Bool
    , totalRewrads : String
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
        |> JP.required "id" JD.int
        |> JP.required "description" JD.string
        |> JP.required "creatorID" JD.int
        |> JP.required "creatorName" JD.string
        |> JP.required "status" JD.int
        |> JP.required "logoFile" JD.string
        |> JP.required "startsAt" JD.string
        |> JP.required "endsAt" JD.string
        |> JP.required "proposals" (JD.list proposalDecoder)
        |> JP.required "supporters" (JD.list supporterDecoder)
        |> JP.required "endsInHours" JD.string
        |> JP.required "endsInMinutes" JD.string
        |> JP.required "isTimeLimit" JD.bool
        |> JP.required "totalRewrads" JD.string


proposalDecoder : JD.Decoder Proposal
proposalDecoder =
    JP.decode Proposal
        |> JP.required "id" JD.int
        |> JP.required "description" JD.string
        |> JP.required "doerId" JD.int
        |> JP.required "doerName" JD.string
        |> JP.required "approved" JD.bool


supporterDecoder : JD.Decoder Supporter
supporterDecoder =
    JP.decode Supporter
        |> JP.required "userId" JD.int
        |> JP.required "userName" JD.string
        |> JP.required "tokenId" JD.int
        |> JP.required "tokenName" JD.string
        |> JP.required "amount" JD.string
        |> JP.required "status" JD.int
