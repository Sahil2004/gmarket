import { Component, Input, OnInit } from '@angular/core';
import { MatInputModule } from '@angular/material/input';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { map, Observable, startWith } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { IStock } from '../../../types';

@Component({
  selector: 'search',
  templateUrl: 'search.html',
  imports: [
    MatFormFieldModule,
    MatInputModule,
    ReactiveFormsModule,
    MatAutocompleteModule,
    MatButtonModule,
    MatIconModule,
    AsyncPipe,
  ],
})
export class Search implements OnInit {
  @Input() title: string = 'Search';
  @Input() options: IStock[] = [];
  @Input() symbol: string = 'search';
  @Input() searchHandler: (symbol: string, exchange: string) => void = (
    symbol: string,
    exchange: string,
  ) => {};

  filteredOptions: Observable<IStock[]> = new Observable<IStock[]>();

  ngOnInit(): void {
    this.filteredOptions = this.search.valueChanges.pipe(
      startWith(''),
      map((value) => this._filter(value || '')),
    );
  }

  search = new FormControl('');

  symbolDeserializer = (value: string): string[] => {
    return value.split(':');
  };

  private _filter(value: string): IStock[] {
    const filterVal = value.toLowerCase();
    return this.options.filter((option) => option.symbol.toLowerCase().includes(filterVal));
  }
}
