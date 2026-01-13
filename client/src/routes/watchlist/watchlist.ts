import { Component } from '@angular/core';
import { RouterLink, RouterOutlet } from '@angular/router';
import { MatSidenavModule } from '@angular/material/sidenav';
import { Search } from '../../components';

@Component({
  selector: 'watchlist',
  templateUrl: 'watchlist.html',
  imports: [RouterLink, RouterOutlet, MatSidenavModule, Search],
})
export class Watchlist {}
