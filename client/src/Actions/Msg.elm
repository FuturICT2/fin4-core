module Actions.Msg exposing (Msg(..))

import Actions.Ports exposing (ImagePortData)
import Common.Json exposing (EmptyResponse)
import Http
import Material
import Model.Actions exposing (Actions)


type Msg
    = Mdl (Material.Msg Msg)
    | OnLoadActionsResponse (Result Http.Error Actions)
    | AddRewards Int String
    | SubmitProposal Int String
    | OnAddRewardsResponse (Result Http.Error EmptyResponse)
    | TickerTimout
    | SetClaim Int String
    | ApproveClaim Int
    | ApproveClaimResponse (Result Http.Error EmptyResponse)
    | OnDoLikeResponse (Result Http.Error EmptyResponse)
    | DoLike Int Bool
    | ImageSelected Int
    | ImageRead ImagePortData
