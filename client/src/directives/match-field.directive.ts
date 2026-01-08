import type { AbstractControl, ValidationErrors, ValidatorFn } from '@angular/forms';

/**
 * This validator checks if the value of one form field matches the value of another form field.
 * @param fieldToMatch The field to match (whose value should be equal to the value of another field, dependent field)
 * @param fieldToMatchWith The field to match with (independent field)
 * @returns fieldsMismatch error if the values of the two fields do not match, null otherwise
 */
export function matchFieldValidator(fieldToMatch: string, fieldToMatchWith: string): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const fieldToMatchControl = control.get(fieldToMatch);
    const fieldToMatchWithControl = control.get(fieldToMatchWith);

    if (!fieldToMatchControl || !fieldToMatchWithControl) {
      return null;
    }
    if (!fieldToMatchControl.touched || !fieldToMatchWithControl.touched) {
      return null;
    }
    if (!fieldToMatchControl.dirty || !fieldToMatchWithControl.dirty) {
      return null;
    }

    return fieldToMatchControl.value !== fieldToMatchWithControl.value
      ? { fieldsMismatch: true }
      : null;
  };
}
