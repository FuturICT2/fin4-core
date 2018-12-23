module Actions.Update exposing (update)

import Actions.Command
    exposing
        ( addRewardsCmd
        , approveProposalCmd
        , loadActionsCmd
        , submitProposalCmd
        )
import Actions.Model exposing (Model)
import Actions.Msg exposing (Msg(..))
import Debug
import Dict exposing (Dict)
import Main.Context exposing (Context)
import Material
import Model.Actions exposing (actionsDecoder)


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        OnLoadActionsResponse resp ->
            case resp of
                Ok actions ->
                    { model | actions = Just actions, error = Nothing } ! []

                Err error ->
                    Debug.log (toString error)
                        { model
                            | actions = Nothing
                            , error = Just "Error loading your actions, please try again!"
                        }
                        ! []

        AddRewards actionId amount ->
            model ! [ addRewardsCmd ctx model actionId amount ]

        OnAddRewardsResponse resp ->
            model ! [ loadActionsCmd ctx 1 ]

        TickerTimout ->
            model ! [ loadActionsCmd ctx 1 ]

        SubmitProposal actionId proposal ->
            let
                props =
                    Dict.remove actionId model.proposals
            in
            { model | proposals = props } ! [ submitProposalCmd ctx model actionId proposal ]

        ApproveProposal proposalId actionId doerId ->
            model ! [ approveProposalCmd ctx model proposalId actionId doerId ]

        ApproveProposalResponse resp ->
            model ! [ loadActionsCmd ctx 1 ]

        SetProposal actionId proposal ->
            let
                props =
                    Dict.insert actionId proposal model.proposals
            in
            { model | proposals = props } ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
