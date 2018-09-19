module Model.Portfolio exposing (Portfolio, Position, init, portfolioDecoder, positionDecoder)

import Json.Decode as JD
import Json.Decode.Pipeline as JP


type alias Position =
    { tokenId : String
    , name : String
    , symbol : String
    , balance : String
    , valueInUSD : String
    }


type alias Portfolio =
    { positions : List Position
    , valueInUSD : String
    }


init : Portfolio
init =
    { positions = []
    , valueInUSD = "0.00"
    }


portfolioDecoder : JD.Decoder Portfolio
portfolioDecoder =
    JP.decode Portfolio
        |> JP.required "Positions" (JD.list positionDecoder)
        |> JP.required "ValueInUSD" JD.string


positionDecoder : JD.Decoder Position
positionDecoder =
    JP.decode Position
        |> JP.required "TokenID" JD.string
        |> JP.required "Name" JD.string
        |> JP.required "Symbol" JD.string
        |> JP.required "Balance" JD.string
        |> JP.required "ValueInUSD" JD.string
