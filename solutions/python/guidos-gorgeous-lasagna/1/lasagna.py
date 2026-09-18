"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language:
https://en.wikipedia.org/wiki/Guido_van_Rossum

This is a module docstring, used to describe the functionality
of a module and its functions and/or classes.
"""

EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2


def bake_time_remaining(elapsed_bake_time):
    """Calculate the bake time remaining.

    Parameters:
        elapsed_bake_time (int): The baking time already elapsed.

    Returns:
        int: The remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.
    """
    return EXPECTED_BAKE_TIME - elapsed_bake_time


def preparation_time_in_minutes(number_of_layers):
    """Calculate the preparation time.

    Parameters:
        number_of_layers (int): The number of layers in the lasagna.

    Returns:
        int: The preparation time (in minutes) derived from 'PREPARATION_TIME'.
    """
    return PREPARATION_TIME * number_of_layers


def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """Calculate the elapsed cooking time.

    Parameters:
        number_of_layers (int): The number of layers in the lasagna.
        elapsed_bake_time (int): The baking time already elapsed.

    Returns:
        int: The total time elapsed (in minutes) since starting to cook,
        which is the sum of preparation time and elapsed bake time.
    """
    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time