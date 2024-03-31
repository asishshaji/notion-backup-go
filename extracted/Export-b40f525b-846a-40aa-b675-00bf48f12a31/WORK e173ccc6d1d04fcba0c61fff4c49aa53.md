# WORK

- Loop through the request
    - see if we have enough quantity to fulfill the order (storage+fulfillment- agg_qty)
        - if you dont have enough qty in fufillment, create a replen demand
            - You calcaulate the replen_qty
            - When creating replen_demand,  check if it had an overflow before
                - if the overflow qty is less than the replen_qty, you subtract the replen_qty from the overflow_qty, why??
                    - Because the overflow represents the qty that was already replened in extra, this can happen when replen in multiples of full_case_capicity, you might over_replen, qties which is stored as overflows
                - And you set the replen_qty to zero
        - 
    - if not short the order