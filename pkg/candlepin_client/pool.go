package candlepin_client

import (
	"context"
	"time"

	caliri "github.com/content-services/caliri/release/v4"
)

func (c *cpClientImpl) FetchPool(ctx context.Context, ownerKey string, productID string) (*caliri.PoolDTO, error) {
	ctx, client, err := getCandlepinClient(ctx)
	if err != nil {
		return nil, err
	}

	found, httpResp, err := client.OwnerAPI.ListOwnerPools(ctx, ownerKey).Product(productID).Execute()
	if httpResp != nil {
		defer httpResp.Body.Close()
	}
	if len(found) == 0 {
		return nil, nil
	}
	if err != nil {
		return nil, errorWithResponseBody("couldn't fetch product", httpResp, err)
	}
	return &found[0], nil
}

func (c *cpClientImpl) CreatePool(ctx context.Context, ownerKey string) (string, error) {
	ctx, client, err := getCandlepinClient(ctx)
	if err != nil {
		return "", err
	}

	productID := GetProductID(ownerKey)
	found, err := c.FetchPool(ctx, ownerKey, productID)
	if found != nil || err != nil {
		return found.GetId(), err
	}

	endDate := time.Date(2049, time.December, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	startDate := time.Now().Format(time.RFC3339)
	quantity := int64(-1)
	found, httpResp, err := client.OwnerAPI.CreatePool(ctx, ownerKey).PoolDTO(caliri.PoolDTO{EndDate: &endDate, StartDate: &startDate, ProductId: &productID, Quantity: &quantity}).Execute()
	if httpResp != nil {
		defer httpResp.Body.Close()
	}
	if err != nil {
		return "", errorWithResponseBody("couldn't create pool", httpResp, err)
	}
	return found.GetId(), nil
}
