module Token.Msg exposing (Msg(..))

import Common.Json exposing (EmptyResponse)
import Http
import Material
import Model.Tokens exposing (Token)
import Token.Ports exposing (ImagePortData)


type Msg
    = Mdl (Material.Msg Msg)
    | OnTokenLoadResponse (Result Http.Error Token)
    | SubmitProposal String
    | SubmitProposalResponse (Result Http.Error EmptyResponse)
    | SetClaim String
    | ApproveClaim Int
    | ApproveClaimResponse (Result Http.Error EmptyResponse)
    | ImageSelected
    | ImageRead ImagePortData
    | OnDoLikeResponse (Result Http.Error EmptyResponse)
    | DoLike Bool
    | SelectTab Int
    | ToggleClaimButton
    | TickerTimout
