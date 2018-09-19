module Common.Date exposing (..)

import Date.Extra as Date


render strDate =
    case Date.fromIsoString strDate of
        Just date ->
            Date.toFormattedString "y-MM-dd HH:mm:ss" date

        Nothing ->
            "invalid:date"
