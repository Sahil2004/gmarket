export interface ITradingAccount {
  cash_balance: number;
  margin_used: number;
  available: number;
  updated_at: string;
}

export interface IBankAccount {
  id: string;
  bank_name: string;
  account_number: string;
  ifsc: string;
  nickname?: string;
  created_at: string;
}

export interface IHolding {
  symbol: string;
  exchange: string;
  product_type: string;
  quantity: number;
  avg_price: number;
  ltp: number;
  last_close: number;
  pnl: number;
  day_pnl: number;
  total_pnl: number;
  current_value: number;
}

export interface IOrder {
  id: string;
  symbol: string;
  exchange: string;
  side: 'buy' | 'sell';
  product_type: 'regular' | 'intraday';
  order_type: 'limit' | 'market';
  quantity: number;
  price: number;
  stop_loss?: number;
  status: 'open' | 'executed' | 'cancelled';
  filled_quantity: number;
  margin_required: number;
  created_at: string;
  executed_at?: string;
}

export interface ITradingSnapshot {
  account: ITradingAccount;
  holdings: IHolding[];
  positions: IHolding[];
  total_holdings_pnl: number;
  total_positions_day_pnl: number;
  open_orders: IOrder[];
  executed_orders: IOrder[];
  bank_accounts: IBankAccount[];
}

export interface IPlaceOrder {
  symbol: string;
  exchange: string;
  side: 'buy' | 'sell';
  product_type: 'regular' | 'intraday';
  order_type: 'limit' | 'market';
  quantity: number;
  price: number;
  stop_loss?: number;
}

export interface IOrderPreview {
  symbol: string;
  exchange: string;
  side: string;
  product_type: string;
  quantity: number;
  price: number;
  ltp: number;
  margin_required: number;
  available: number;
  order_value: number;
}

export interface IOrderDialogData {
  symbol: string;
  exchange: string;
  side: 'buy' | 'sell';
  ltp?: number;
}
