import { Component, Input, output } from '@angular/core';
import { MatTabsModule } from '@angular/material/tabs';
import { StocksList } from '../stocks-list/stocks-list';

@Component({
  selector: 'watchlist-tabs',
  templateUrl: 'watchlist-tabs.html',
  imports: [MatTabsModule, StocksList],
})
export class WatchlistTabs {
  @Input() tabChangeHandler: (event: { index: number }) => void = (event) => {};
  @Input({ required: true }) watchlists!: string[][];
}
