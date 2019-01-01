module Actions.Command exposing
    ( addRewardsCmd
    , approveProposalCmd
    , commands
    , likeCmd
    , loadActionsCmd
    , submitProposalCmd
    )

import Actions.Model exposing (Model)
import Actions.Msg exposing (Msg(..))
import Common.Api exposing (get, postWithCsrf)
import Common.Json exposing (decodeAt, deocdeIntWithDefault, emptyResponseDecoder)
import Dict
import Json.Decode as JD
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import Model.Actions exposing (actionsDecoder)


commands : Context -> Cmd Msg
commands ctx =
    case ctx.route of
        ActionsRoute ->
            loadActionsCmd ctx 1

        _ ->
            Cmd.none


addRewardsCmd : Context -> Model -> Int -> String -> Cmd Msg
addRewardsCmd ctx model actionId amount =
    postWithCsrf ctx
        OnAddRewardsResponse
        "/add-rewards-to-action"
        (encodeAddRewardsAction model actionId amount)
        emptyResponseDecoder


approveProposalCmd : Context -> Model -> Int -> Cmd Msg
approveProposalCmd ctx model claimId =
    postWithCsrf ctx
        ApproveClaimResponse
        "/approve-proposal"
        (encodeApproveProposal claimId)
        emptyResponseDecoder


submitProposalCmd : Context -> Model -> Int -> String -> Cmd Msg
submitProposalCmd ctx model actionId proposal =
    let
        v =
            Maybe.withDefault
                { contents = ""
                , filename = ""
                }
            <|
                Dict.get actionId model.claimImages
    in
    postWithCsrf ctx
        OnAddRewardsResponse
        "/submit-proposal"
        (encodeSubmitProposal actionId proposal v.contents)
        emptyResponseDecoder


likeCmd ctx tokenId state =
    let
        url =
            case state of
                True ->
                    "/unlike/" ++ toString tokenId

                False ->
                    "/like/" ++ toString tokenId
    in
    get ctx
        OnDoLikeResponse
        url
        []
        emptyResponseDecoder


encodeAddRewardsAction : Model -> Int -> String -> JE.Value
encodeAddRewardsAction model actionId amount =
    JE.object
        [ ( "actionId", JE.int actionId )
        , ( "tokenId", JE.int 1 )
        , ( "amount", JE.string amount )
        ]


encodeApproveProposal : Int -> JE.Value
encodeApproveProposal claimId =
    JE.object
        [ ( "claimId", JE.int claimId )
        ]


encodeSubmitProposal tokenId claim base64 =
    JE.object
        [ ( "tokenId", JE.int tokenId )
        , ( "proposal", JE.string claim )
        , ( "image64", JE.string base64 )
        ]


loadActionsCmd ctx page =
    get ctx
        OnLoadActionsResponse
        "/actions"
        []
        actionsDecoder
