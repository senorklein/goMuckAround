#!/usr/bin/python

import argparse

def dostuff(i):

    print("input: " + i)



def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("-i", "--input",  action="store")
    args = parser.parse_args()
    dostuff(args.input)



if __name__  == "__main__":
    main()