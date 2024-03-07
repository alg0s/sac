WITH monthly_orders AS (
    SELECT 
        DATE_TRUNC('month', o.ordered_at) AS month
        , c.customer_id
        , SUM(oli.unit_price * oli.quantity) AS total_monthly_order_value
    FROM 
        customers c
        INNER JOIN orders o ON c.customer_id = o.customer_id
        INNER JOIN order_line_items oli ON o.order_id = oli.order_id
    GROUP BY 
        DATE_TRUNC('month', o.ordered_at)
        , c.customer_id
),
ranked_customers AS (
    SELECT 
        EXTRACT(YEAR FROM month) AS year
        , EXTRACT(MONTH FROM month) AS month
        , customer_id
        , total_monthly_order_value
        , ROW_NUMBER() OVER(
            PARTITION BY month 
            ORDER BY 
                total_monthly_order_value DESC
                , customer_id
            ) AS rank
    FROM 
        monthly_orders
)
SELECT 
    year
    , month
    , customer_id
    , total_monthly_order_value
FROM 
    ranked_customers
WHERE 
    rank = 1
ORDER BY 
    YEAR
    , MONTH
;