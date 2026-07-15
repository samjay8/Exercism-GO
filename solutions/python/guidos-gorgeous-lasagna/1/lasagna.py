"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language:
https://en.wikipedia.org/wiki/Guido_van_Rossum

This is a module docstring, used to describe the functionality
of a module and its functions and/or classes.
"""


#TODO (student): define your EXPECTED_BAKE_TIME (required) and PREPARATION_TIME (optional) constants below.
EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2

#TODO (student): Remove 'pass' and complete the 'bake_time_remaining()' function below.
def bake_time_remaining(actual_minutes):
    """ This is a function that returns the remaining minutes the lasagna needs to bake"""
    return EXPECTED_BAKE_TIME - actual_minutes

#TODO (student): Define the 'preparation_time_in_minutes()' function below.
def preparation_time_in_minutes(number_of_layers):
    """This is a function that returns the preparation time needed to make them in minutes"""
    return number_of_layers * PREPARATION_TIME
    

#TODO (student): define the 'elapsed_time_in_minutes()' function below.
def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """ This is a function that returns the total elapsed times you have been cooking."""
    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time

# TODO (student): Remember to go back and add docstrings to all your functions
