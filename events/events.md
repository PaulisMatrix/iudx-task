# WebPage Events
## what type of events to capture

Example:
```
{
  "timestamp": "2023-09-25T14:30:00",
  "userId": "user123",
  "eventType": "pageView",
  "productId": "product456",
  "sessionDuration": 180
}
```

1. **PageView** Event: UserID, EventType, CreatedAt, Duration, PageURL, ProductID

2. **Click** Event: UserID, EventType, CreatedAt, ClickedTarget, ProductID

3. **Search** Event: UserID, EventType, CreatedAt, Duration, SearchQuery, PageURL

4. **Purchase** Event: UserID, EventType, CreatedAt, Duration, PaymentMode, NumItemsPurchased

5. **AdEvent** Event: UserID, EventType, CreatedAt, Duration, AdvertID, SourcePageURL, DestPageURL

6. **ScrollEvent** Event: UserID, EventType, CreatedAt, Duration, PageURL, ScrolledNumProducts


