module Actions.Msg exposing (Msg(..))

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
    | SetProposal Int String
    | ApproveProposal Int Int Int
    | ApproveProposalResponse (Result Http.Error EmptyResponse)
